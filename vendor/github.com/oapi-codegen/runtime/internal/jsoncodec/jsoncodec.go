// Package jsoncodec holds the process-wide JSON codec that the runtime package
// and runtime/types share.
//
// It is internal because runtime imports types, so types cannot import
// runtime. The runtime package re-exports its symbols, as aliases of the types
// and one-line wrappers of the functions, so that nothing is defined twice.
package jsoncodec

import (
	"bytes"
	"encoding/json"
	"io"
)

// Codec encodes and decodes JSON. runtime.JSONCodec is an alias of it.
type Codec interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

// Streaming is optionally implemented by a Codec that encodes to and decodes
// from a stream. runtime.StreamingJSONCodec is an alias of it.
type Streaming interface {
	MarshalWrite(w io.Writer, v any) error
	UnmarshalRead(r io.Reader, v any) error
}

// Std is encoding/json's json.Marshal and json.Unmarshal.
// runtime.StdJSONCodec is an alias of it.
//
// It deliberately doesn't implement Streaming: a codec that embeds it and
// overrides Marshal or Unmarshal then encodes and decodes streams through
// its own methods, via MarshalWrite and UnmarshalRead's fall-through, rather
// than through encoding/json's Encoder and Decoder promoted from Std.
type Std struct{}

func (Std) Marshal(v any) ([]byte, error)      { return json.Marshal(v) }
func (Std) Unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }

// std is the default codec: Std, plus encoding/json's Encoder and Decoder for
// streams, so that the default behaviour is encoding/json's throughout.
type std struct{ Std }

func (std) MarshalWrite(w io.Writer, v any) error  { return json.NewEncoder(w).Encode(v) }
func (std) UnmarshalRead(r io.Reader, v any) error { return json.NewDecoder(r).Decode(v) }

var _ Streaming = std{}

var codec Codec = std{}

// Set replaces the codec. nil restores the default, and so does Std itself:
// it is the default's byte half, and passing it is the natural way to say
// "encoding/json", not a request for the buffered stream path.
func Set(c Codec) {
	if c == nil || c == (Std{}) {
		c = std{}
	}
	codec = c
}

// Marshal encodes v with the codec.
func Marshal(v any) ([]byte, error) { return codec.Marshal(v) }

// Unmarshal decodes data into v with the codec.
func Unmarshal(data []byte, v any) error { return codec.Unmarshal(data, v) }

// MarshalWrite encodes v to w: with the codec's MarshalWrite when it
// implements Streaming, otherwise with Marshal followed by a single Write.
func MarshalWrite(w io.Writer, v any) error {
	if s, ok := codec.(Streaming); ok {
		return s.MarshalWrite(w, v)
	}
	data, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// UnmarshalRead decodes v from r: with the codec's UnmarshalRead when it
// implements Streaming, otherwise by reading r to its end and calling
// Unmarshal. When r is nil or holds no value, nothing or only JSON white
// space, it returns io.EOF without involving the codec, whose own report of
// an empty input varies.
func UnmarshalRead(r io.Reader, v any) error {
	r, err := nonEmpty(r)
	if err != nil {
		return err
	}
	if s, ok := codec.(Streaming); ok {
		return s.UnmarshalRead(r, v)
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return codec.Unmarshal(data, v)
}

// nonEmpty returns io.EOF when r is nil or has only JSON white space before
// its end, and otherwise r with the bytes read to find that out replayed in
// front of it, so that the codec sees the input unchanged, a syntax error's
// offset included.
func nonEmpty(r io.Reader) (io.Reader, error) {
	if r == nil {
		return nil, io.EOF
	}
	var read []byte
	for {
		var b [1]byte
		n, err := r.Read(b[:])
		if n == 1 {
			read = append(read, b[0])
			switch b[0] {
			case ' ', '\t', '\r', '\n':
				continue
			}
			return io.MultiReader(bytes.NewReader(read), r), nil
		}
		if err != nil {
			return nil, err
		}
	}
}
