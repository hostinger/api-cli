// Package logging builds the connector logger and copies warnings into the
// error ring that `status` shows.
package logging

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/hostinger/hcdn-connector/internal/session"
	"github.com/hostinger/hcdn-connector/internal/termsafe"
	"github.com/hostinger/hcdn-connector/pkg/config"
)

func New(w io.Writer, cfg config.Log, ring *session.ErrorRing) *slog.Logger {
	var lvl slog.Level
	switch cfg.Level {
	case "debug":
		lvl = slog.LevelDebug
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: lvl}
	var h slog.Handler
	if cfg.Format == "text" {
		// The text handler quotes C1 and invalid UTF-8 but treats DEL as safe.
		h = slog.NewTextHandler(escWriter{w, escapeDEL}, opts)
	} else {
		h = slog.NewJSONHandler(escWriter{w, termsafe.JSON}, opts)
	}
	return slog.New(Tee(h, ring))
}

// escWriter applies esc to each Write. The slog handlers write a whole record
// in one Write, so a multi-byte rune is never split across calls.
type escWriter struct {
	w   io.Writer
	esc func([]byte) []byte
}

func (e escWriter) Write(p []byte) (int, error) {
	if _, err := e.w.Write(e.esc(p)); err != nil {
		return 0, err
	}
	return len(p), nil
}

// escapeDEL rewrites DEL as \x7f, matching how the text handler quotes it.
func escapeDEL(b []byte) []byte {
	if bytes.IndexByte(b, 0x7f) < 0 {
		return b
	}
	return bytes.ReplaceAll(b, []byte{0x7f}, []byte(`\x7f`))
}

// Tee copies Warn+ records that h accepts into ring.
func Tee(h slog.Handler, ring *session.ErrorRing) slog.Handler {
	return &teeHandler{inner: h, ring: ring}
}

type teeHandler struct {
	inner slog.Handler
	ring  *session.ErrorRing
	attrs []slog.Attr
}

func (t *teeHandler) Enabled(ctx context.Context, l slog.Level) bool { return t.inner.Enabled(ctx, l) }

func (t *teeHandler) Handle(ctx context.Context, r slog.Record) error {
	if t.ring != nil && r.Level >= slog.LevelWarn {
		msg := r.Message
		r.Attrs(func(a slog.Attr) bool {
			msg += " " + a.Key + "=" + a.Value.String()
			return true
		})
		for _, a := range t.attrs {
			msg += " " + a.Key + "=" + a.Value.String()
		}
		t.ring.Add(session.ErrorEntry{Time: r.Time.UTC().Truncate(time.Second), Level: r.Level.String(), Message: msg})
	}
	return t.inner.Handle(ctx, r)
}

func (t *teeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &teeHandler{inner: t.inner.WithAttrs(attrs), ring: t.ring, attrs: append(append([]slog.Attr{}, t.attrs...), attrs...)}
}

func (t *teeHandler) WithGroup(name string) slog.Handler {
	return &teeHandler{inner: t.inner.WithGroup(name), ring: t.ring, attrs: t.attrs}
}
