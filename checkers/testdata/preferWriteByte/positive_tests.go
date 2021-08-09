package checker_test

import (
	"strings"
)

func _(b *strings.Builder) {
	/*! consider replacing b.WriteRune('\n') with b.WriteByte('\n') */
	b.WriteRune('\n')
}
