package runtime

import (
	"io"

	"github.com/oapi-codegen/runtime/internal/jsoncodec"
)

// JSONCodec encodes and decodes JSON for this package and for the code
// oapi-codegen generates. It has two methods:
//
//	Marshal(v any) ([]byte, error)
//	Unmarshal(data []byte, v any) error
//
// A codec must produce and accept standard JSON: it selects an implementation
// and its options, such as encoding/json/v2 or an encoder that doesn't escape
// HTML, not a different format. A codec that also implements
// StreamingJSONCodec encodes and decodes request and response bodies as
// streams.
//
// A codec must be re-entrant. The MarshalJSON and UnmarshalJSON methods of
// the types in runtime/types, and of the union and additionalProperties
// types oapi-codegen generates, call back into the codec from inside the
// codec's own Marshal or Unmarshal of a value that contains them. A codec
// that holds a lock across Marshal deadlocks on the first such value.
type JSONCodec = jsoncodec.Codec

// StreamingJSONCodec is optionally implemented by a JSONCodec that encodes to
// and decodes from a stream, which request and response bodies use:
//
//	MarshalWrite(w io.Writer, v any) error
//	UnmarshalRead(r io.Reader, v any) error
//
// Without it, JSONMarshalWrite and JSONUnmarshalRead buffer the whole body
// through Marshal and Unmarshal. Whether MarshalWrite ends the value with a
// newline, and whether UnmarshalRead rejects data after the value, is up to
// the implementation: encoding/json's Encoder and Decoder do and don't,
// respectively. Neither this package nor generated code relies on either.
// UnmarshalRead is never handed an empty input: see JSONUnmarshalRead.
type StreamingJSONCodec = jsoncodec.Streaming

// StdJSONCodec is encoding/json's json.Marshal and json.Unmarshal. Embed it in
// a codec to replace one method.
//
// It doesn't implement StreamingJSONCodec, so a codec that embeds it handles
// bodies through its own Marshal and Unmarshal, buffered. That differs from
// the default in two ways: a response body has no trailing newline, and data
// after a request body's value is rejected where encoding/json's Decoder
// ignores it. SetJSONCodec treats StdJSONCodec itself as the default, so
// passing it doesn't select the buffered path.
type StdJSONCodec = jsoncodec.Std

// SetJSONCodec replaces the codec used by this package and by the code
// oapi-codegen generates, process-wide. It is for a program's main package to
// call during initialisation, not for a library: it is not synchronised, and
// it changes the JSON of every generated package linked into the program,
// including those of other libraries. nil restores the default, encoding/json
// throughout, and so does StdJSONCodec.
//
// The codec covers what this package and generated code encode and decode
// themselves. It doesn't reach JSON that a web framework handles on their
// behalf: the strict servers for echo, gin, fiber and iris decode request
// bodies with the framework's binder, fiber's and iris's write responses with
// the framework's JSON method, and a request body a validation middleware
// decoded and kept (see BindJSONBody) is bound without decoding it again.
func SetJSONCodec(c JSONCodec) { jsoncodec.Set(c) }

// JSONMarshal encodes v with the current JSONCodec.
func JSONMarshal(v any) ([]byte, error) { return jsoncodec.Marshal(v) }

// JSONUnmarshal decodes data into v with the current JSONCodec.
func JSONUnmarshal(data []byte, v any) error { return jsoncodec.Unmarshal(data, v) }

// JSONMarshalWrite encodes v to w with the current JSONCodec: with its
// MarshalWrite when it implements StreamingJSONCodec, otherwise with Marshal
// followed by a single Write.
func JSONMarshalWrite(w io.Writer, v any) error { return jsoncodec.MarshalWrite(w, v) }

// JSONUnmarshalRead decodes v from r with the current JSONCodec: with its
// UnmarshalRead when it implements StreamingJSONCodec, otherwise by reading r
// to its end and calling Unmarshal.
//
// When r is nil or holds no value, nothing or only JSON white space, it
// returns io.EOF whichever codec is in use, as encoding/json's Decoder does,
// and the codec isn't involved: generated handlers take io.EOF as an optional
// request body left out. Otherwise the codec sees r's input unchanged.
func JSONUnmarshalRead(r io.Reader, v any) error { return jsoncodec.UnmarshalRead(r, v) }
