// Package termsafe neutralises control characters in text that came from a
// network peer before it reaches a terminal.
package termsafe

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// String escapes C0/C1 control characters and DEL as \xNN so peer-supplied
// text cannot drive the terminal; invalid UTF-8 bytes become U+FFFD.
func String(s string) string {
	if utf8.ValidString(s) && !strings.ContainsFunc(s, isControl) {
		return s
	}
	var b strings.Builder
	for _, r := range s {
		if isControl(r) {
			fmt.Fprintf(&b, `\x%02x`, r)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// JSON escapes DEL and C1 control characters in encoded JSON as \u00XX, and
// invalid UTF-8 bytes as �. encoding/json and log/slog escape only C0,
// which leaves DEL and the C1 range (CSI, OSC, DCS) raw. Outside strings JSON
// is plain ASCII, so every rune this rewrites sits inside a string, where the
// escape decodes back to the same value.
func JSON(b []byte) []byte {
	i := 0
	for i < len(b) {
		r, n := utf8.DecodeRune(b[i:])
		if (r == utf8.RuneError && n == 1) || (r >= 0x7f && r <= 0x9f) {
			break
		}
		i += n
	}
	if i == len(b) {
		return b
	}
	out := append(make([]byte, 0, len(b)+16), b[:i]...)
	for i < len(b) {
		r, n := utf8.DecodeRune(b[i:])
		switch {
		case r == utf8.RuneError && n == 1:
			out = append(out, `�`...)
		case r >= 0x7f && r <= 0x9f:
			out = fmt.Appendf(out, `\u%04x`, r)
		default:
			out = append(out, b[i:i+n]...)
		}
		i += n
	}
	return out
}

// Printable reports whether s is valid UTF-8 made only of runes for which
// unicode.IsPrint holds. Peer-supplied identifiers are checked with it.
func Printable(s string) bool {
	return utf8.ValidString(s) && !strings.ContainsFunc(s, func(r rune) bool { return !unicode.IsPrint(r) })
}

func isControl(r rune) bool { return r < 0x20 || (r >= 0x7f && r <= 0x9f) }
