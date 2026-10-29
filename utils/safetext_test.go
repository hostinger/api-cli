package utils

import (
	"errors"
	"testing"
)

func TestSafeText(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"CSI", "a\x1b[2Kb", "a\uFFFD[2Kb"},
		{"OSC 52 BEL", "\x1b]52;c;aGk=\x07", "\uFFFD]52;c;aGk=\uFFFD"},
		{"OSC 52 ST", "\x1b]52;c;aGk=\x1b\\", "\uFFFD]52;c;aGk=\uFFFD\\"},
		{"CR", "ok\rfake", "ok\uFFFDfake"},
		{"BS", "ab\bc", "ab\uFFFDc"},
		{"DEL", "a\x7fb", "a\uFFFDb"},
		{"C1 CSI", "a\u009b2Kb", "a\uFFFD2Kb"},
		{"UTF-8", "Café naïve 🚀 日本", "Café naïve 🚀 日本"},
		{"newline and tab", "a\nb\tc", "a\nb\tc"},
		{"empty", "", ""},
	}
	for _, c := range cases {
		if got := SafeText(c.in); got != c.want {
			t.Errorf("%s: SafeText(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}

func TestSafeLog(t *testing.T) {
	cases := []struct{ name, in, colour, plain string }{
		{"SGR", "\x1b[31merr\x1b[0m", "\x1b[31merr\x1b[0m", "err"},
		{"SGR colon", "\x1b[38:5:1mx\x1b[m", "\x1b[38:5:1mx\x1b[m", "x"},
		{"OSC 0 BEL", "a\x1b]0;pwned\x07b", "ab", "ab"},
		{"OSC 8 ST", "\x1b]8;;https://evil\x1b\\link\x1b]8;;\x1b\\", "link", "link"},
		{"OSC 52 BEL", "\x1b]52;c;aGk=\x07ok", "ok", "ok"},
		{"OSC 52 ST", "\x1b]52;c;aGk=\x1b\\ok", "ok", "ok"},
		{"erase screen", "a\x1b[2Jb", "ab", "ab"},
		{"cursor up", "a\x1b[1Ab", "ab", "ab"},
		{"private mode", "\x1b[?25lx", "x", "x"},
		{"unterminated OSC", "\x1b]0;t", "\uFFFD]0;t", "\uFFFD]0;t"},
		{"CR", "ok\rfake", "ok\nfake", "ok\nfake"},
		{"CRLF", "a\r\nb\r\n", "a\nb\n", "a\nb\n"},
		{"progress bar", "10%\r50%\r100%\r\ndone", "10%\n50%\n100%\ndone", "10%\n50%\n100%\ndone"},
		{"coloured CRLF", "\x1b[32mok\x1b[0m\r\n", "\x1b[32mok\x1b[0m\n", "ok\n"},
		{"BS", "ab\bc", "ab\uFFFDc", "ab\uFFFDc"},
		{"C1 CSI", "a\u009b2Jb", "a\uFFFD2Jb", "a\uFFFD2Jb"},
		{"DEL", "a\x7fb", "a\uFFFDb", "a\uFFFDb"},
		{"newline tab UTF-8", "Café\n\t🚀 日本", "Café\n\t🚀 日本", "Café\n\t🚀 日本"},
		{"empty", "", "", ""},
	}
	for _, c := range cases {
		if got := SafeLog(c.in, true); got != c.colour {
			t.Errorf("%s: SafeLog(%q, true) = %q, want %q", c.name, c.in, got, c.colour)
		}
		if got := SafeLog(c.in, false); got != c.plain {
			t.Errorf("%s: SafeLog(%q, false) = %q, want %q", c.name, c.in, got, c.plain)
		}
	}
}

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, errors.New("closed") }

func TestSafeWriterReturnsError(t *testing.T) {
	if n, err := (safeWriter{failWriter{}}).Write([]byte("x")); err == nil || n != 0 {
		t.Errorf("Write = %d, %v; want 0 and an error", n, err)
	}
}
