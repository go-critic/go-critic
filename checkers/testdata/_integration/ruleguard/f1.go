package foo

import "fmt"

type myError error

func _() {
	var s1, s2 string
	var _ = fmt.Sprintf("%s%s", s1, s2)
}

func _() {
	var s1, s2 string
	if s1 == s2 {
		if s1 != "" {
			println(s1, s2)
		}
	}
}
