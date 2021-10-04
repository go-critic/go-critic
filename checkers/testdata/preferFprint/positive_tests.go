package checker_test

import (
	"fmt"
	"io"
)

func _(w io.Writer) {
	/*! fmt.Fprintf(w, "%x", 10) should be preferred */
	w.Write([]byte(fmt.Sprintf("%x", 10)))
	/*! fmt.Fprint(w, 1, 2, 3, 4) should be preferred */
	w.Write([]byte(fmt.Sprint(1, 2, 3, 4)))
	/*! fmt.Fprintln(w, 1, 2, 3, 4) should be preferred */
	w.Write([]byte(fmt.Sprintln(1, 2, 3, 4)))
}
