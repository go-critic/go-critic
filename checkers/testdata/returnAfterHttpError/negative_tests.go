package checker_test

import "net/http"

func _(w http.ResponseWriter, err error) {
	if err != nil {
		foo()
		http.Error(w, "err", 503)
		return
	}
}

func _(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "err", 503)
		return
	}

	foo()
}

func _(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "err", 503)
	}
}

func _(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "err", 503)
	}

	// This is a comment, not a statement.
}

func _(w http.ResponseWriter, err error) {
	func() {
		http.Error(w, "err", 503)
	}()
}
