package checker_test

import (
	"context"
	"net/http"
	"net/http/httptest"
)

func badCases() {
	/*! http.NoBody should be preferred to the nil request body */
	_, _ = http.NewRequest("GET", "https://some.url.com/", nil)

	/*! http.NoBody should be preferred to the nil request body */
	_, _ = http.NewRequestWithContext(context.TODO(), "GET", "https://some.url.com/", nil)

	/*! http.NoBody should be preferred to the nil request body */
	_ = httptest.NewRequest("GET", "https://some.url.com/", nil)
}
