package linttest

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var (
	warningDirectiveRE = regexp.MustCompile(`^\s*/\*! (.*) \*/`)
)

type warnings map[int][]string

func newWarnings(r io.Reader) (warnings, error) {
	ws := make(warnings)
	var pending []string

	s := bufio.NewScanner(r)
	for i := 0; s.Scan(); i++ {
		if m := warningDirectiveRE.FindStringSubmatch(s.Text()); m != nil {
			pending = append(pending, m[1])
		} else if len(pending) != 0 {
			line := i + 1
			ws[line] = pending
			pending = nil
		}
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("read test file data: %w", err)
	}

	return ws, nil
}

func (ws warnings) find(line int, text string) *string {
	for i := range ws[line] {
		if text == ws[line][i] {
			return &ws[line][i]
		}
	}
	return nil
}
