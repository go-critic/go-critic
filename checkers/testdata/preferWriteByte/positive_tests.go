package checker_test

import (
	"strings"
)

func _(b *strings.Builder) {
	b.WriteRune('ь')
	b.WriteByte('\n')
}
