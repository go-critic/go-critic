package foo

import (
	"fmt"
	"regexp"
)

type myError error

func _() {
	var s1, s2 string
	r, _ := regexp.MatchString(`^\s`, fmt.Sprintf("%s%s", s1, s2))
	println(r)
}

func _() {
	var s1, s2 string
	if s1 == s2 {
		if s1 != "" {
			println(s1, s2)
		}
	}
}

func _() {
	k := func() string {
		return "test"
	}
	_ = fmt.Errorf(k())
}
