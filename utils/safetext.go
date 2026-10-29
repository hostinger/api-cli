package utils

import (
	"io"
	"regexp"
	"strings"
	"unicode"
)

// SafeText replaces control characters other than \n and \t (C0 including
// ESC, BEL, CR and BS; DEL; C1 such as U+009B CSI) with U+FFFD, so remote
// text printed to a terminal cannot start escape sequences. Callers that must
// keep some sequences (e.g. SGR colours in build logs) can split those out and
// pass the text between them through SafeText.
func SafeText(s string) string {
	return strings.Map(func(r rune) rune {
		if r != '\n' && r != '\t' && unicode.IsControl(r) {
			return unicode.ReplacementChar
		}
		return r
	}, s)
}

var (
	escSeq = regexp.MustCompile(`\x1b\[[0-?]*[ -/]*[@-~]|\x1b\][^\x07\x1b]*(?:\x07|\x1b\\)`)
	sgrSeq = regexp.MustCompile(`^\x1b\[[0-9;:]*m$`)
	// A lone \r (progress bars) starts a new line rather than overwriting it.
	crlf = strings.NewReplacer("\r\n", "\n", "\r", "\n")
)

// SafeLog makes remote build logs safe for a terminal: complete CSI and OSC
// sequences are removed, except SGR colours when keepColour is set, and the
// rest goes through SafeText. Carriage returns become line breaks.
func SafeLog(s string, keepColour bool) string {
	s = crlf.Replace(s)
	var b strings.Builder
	last := 0
	for _, m := range escSeq.FindAllStringIndex(s, -1) {
		b.WriteString(SafeText(s[last:m[0]]))
		if seq := s[m[0]:m[1]]; keepColour && sgrSeq.MatchString(seq) {
			b.WriteString(seq)
		}
		last = m[1]
	}
	b.WriteString(SafeText(s[last:]))
	return b.String()
}

// safeWriter passes everything written through SafeText.
type safeWriter struct{ w io.Writer }

func (s safeWriter) Write(p []byte) (int, error) {
	if _, err := io.WriteString(s.w, SafeText(string(p))); err != nil {
		return 0, err
	}
	return len(p), nil
}
