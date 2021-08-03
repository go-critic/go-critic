package checker_test

import (
	"strings"
)

func _(b *strings.Builder) {
	b.WriteRune('ь')
	b.WriteByte('\n')
}

type RuneWriter interface {
	WriteRune(r rune) (int, error)
}

func notByteWriter(w RuneWriter) {
	w.WriteRune('\n')
}
