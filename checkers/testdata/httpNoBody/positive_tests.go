package checker_test

import "net/http"

func badCases() {
	/*! http.NoBody should be preferred to the nil request body */
	_, _ = http.NewRequest("GET", "https://some.url.com/", nil)
}
