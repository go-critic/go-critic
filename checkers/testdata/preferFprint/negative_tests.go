package checker_test

import (
	"fmt"
	"io"
)

type f struct{}

func (f) Sprintf(args ...interface{}) string  { return "abc" }
func (f) Sprint(args ...interface{}) string   { return "abc" }
func (f) Sprintln(args ...interface{}) string { return "abc" }

func _(w io.Writer) {
	var fmt f
	w.Write([]byte(fmt.Sprintf()))
	w.Write([]byte(fmt.Sprint()))
	w.Write([]byte(fmt.Sprintln()))

	w.Write([]byte(fmt.Sprintf("%d", 1)))
	w.Write([]byte(fmt.Sprint(1)))
	w.Write([]byte(fmt.Sprintln(1, 2)))
}

func _(w io.StringWriter) {
	w.WriteString(fmt.Sprintf("%x", 10))
	w.WriteString(fmt.Sprint(1, 2, 3, 4))
	w.WriteString(fmt.Sprintln(1, 2, 3, 4))
}

type falseWriteString struct{}

func (f falseWriteString) WriteString(s string) {}

func _(w falseWriteString) {
	w.WriteString(fmt.Sprintf("%x", 10))
	w.WriteString(fmt.Sprint(1, 2, 3, 4))
	w.WriteString(fmt.Sprintln(1, 2, 3, 4))
}
