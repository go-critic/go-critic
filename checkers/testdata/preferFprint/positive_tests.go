package checker_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func _(w io.Writer) {
	/*! fmt.Fprintf(w, "%x", 10) should be preferred to the w.Write([]byte(fmt.Sprintf("%x", 10))) */
	w.Write([]byte(fmt.Sprintf("%x", 10)))
	/*! fmt.Fprint(w, 1, 2, 3, 4) should be preferred to the w.Write([]byte(fmt.Sprint(1, 2, 3, 4))) */
	w.Write([]byte(fmt.Sprint(1, 2, 3, 4)))
	/*! fmt.Fprintln(w, 1, 2, 3, 4) should be preferred to the w.Write([]byte(fmt.Sprintln(1, 2, 3, 4))) */
	w.Write([]byte(fmt.Sprintln(1, 2, 3, 4)))

	buf := &bytes.Buffer{}
	/*! fmt.Fprintf(buf, "%x", 10) should be preferred to the buf.Write([]byte(fmt.Sprintf("%x", 10))) */
	buf.Write([]byte(fmt.Sprintf("%x", 10)))
	/*! fmt.Fprint(buf, 1, 2, 3, 4) should be preferred to the buf.Write([]byte(fmt.Sprint(1, 2, 3, 4))) */
	buf.Write([]byte(fmt.Sprint(1, 2, 3, 4)))
	/*! fmt.Fprintln(buf, 1, 2, 3, 4) should be preferred to the buf.Write([]byte(fmt.Sprintln(1, 2, 3, 4))) */
	buf.Write([]byte(fmt.Sprintln(1, 2, 3, 4)))

	var i int

	/*! suggestion: fmt.Fprint(buf, i) */
	io.WriteString(buf, fmt.Sprint(i))

	/*! suggestion: fmt.Fprintf(buf, "<%4d>", i) */
	io.WriteString(buf, fmt.Sprintf("<%4d>", i))

	/*! suggestion: fmt.Fprintln(buf, i, i) */
	io.WriteString(buf, fmt.Sprintln(i, i))

	/*! suggestion: fmt.Fprint(os.Stdout, i, i) */
	io.WriteString(os.Stdout, fmt.Sprint(i, i))

	/*! suggestion: fmt.Fprint(buf, i) */
	buf.WriteString(fmt.Sprint(i))

	/*! suggestion: fmt.Fprintf(buf, "<%4d>", i) */
	buf.WriteString(fmt.Sprintf("<%4d>", i))

	/*! suggestion: fmt.Fprintln(buf, i, i) */
	buf.WriteString(fmt.Sprintln(i, i))

	/*! suggestion: fmt.Fprint(os.Stdout, i, i) */
	os.Stdout.WriteString(fmt.Sprint(i, i))
}
