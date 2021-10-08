package checker_test

import (
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
