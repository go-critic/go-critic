package checker_test

import "net/http"

func _(w http.ResponseWriter, err error) {
	if err != nil {
		foo()
		http.Error(w, "err", 503)
		return
	}
}
