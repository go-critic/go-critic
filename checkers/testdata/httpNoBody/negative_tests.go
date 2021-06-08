package checker_test

import "net/http"

func goodCases() {
	_, _ = http.NewRequest("GET", "https://some.url.com/", http.NoBody)
}
