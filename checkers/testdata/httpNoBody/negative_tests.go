package checker_test

import (
	"context"
	"net/http"
	"net/http/httptest"
)

func goodCases() {
	_, _ = http.NewRequest("GET", "https://some.url.com/", http.NoBody)
	_, _ = http.NewRequestWithContext(context.TODO(), "GET", "https://some.url.com/", http.NoBody)
	_ = httptest.NewRequest("GET", "https://some.url.com/", http.NoBody)
}
