package checker_test

import (
	"bytes"
	"net/http"
	"strings"
	"sync"
)

func appliedSuggestions(s string, b []byte) {
	var wg sync.WaitGroup
	wg.Done()

	strings.ToTitle(s)
	strings.ReplaceAll(s, "a", "b")

	bytes.ToUpper(b)
	bytes.ToLower(b)
	bytes.ToTitle(b)

	bytes.ReplaceAll(b, b, b)

	_ = http.NotFoundHandler()
}

func nonMatchingArgs(s string, b []byte) {
	var wg sync.WaitGroup
	wg.Add(1)

	strings.Map(nil, s)

	strings.Replace(s, "a", "b", 1)

	bytes.Map(nil, b)
	bytes.Map(nil, b)
	bytes.Map(nil, b)

	bytes.Replace(b, b, b, 1)

	_ = http.HandlerFunc(nil)
}
