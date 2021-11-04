package checker_test

import (
	"strings"
)

func _(b *strings.Builder) {
	/*! consider writing single byte rune '\n' with b.WriteByte('\n') */
	b.WriteRune('\n')
}
