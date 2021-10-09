package checker_test

import (
	"context"
	"net/http"
)

func badCases() {
	/*! http.NoBody should be preferred to the nil request body */
	_, _ = http.NewRequest("GET", "https://some.url.com/", nil)

	/*! http.NoBody should be preferred to the nil request body */
	_, _ = http.NewRequestWithContext(context.TODO(), "GET", "https://some.url.com/", nil)
}
