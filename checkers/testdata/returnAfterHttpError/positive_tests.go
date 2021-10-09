package checker_test

import "net/http"

func _(w http.ResponseWriter, err error) {
	if err != nil {
		foo()
		/*! Possibly return is missed after the http.Error call */
		http.Error(w, "err", 503)
	}

	if err != nil {
		/*! Possibly return is missed after the http.Error call */
		http.Error(w, "err", 505)
	}
}

func foo() {}
