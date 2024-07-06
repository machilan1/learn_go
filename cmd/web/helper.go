package main

import "net/http"

// Custom error wrapper for server errors
func (a *application) serverError(w http.ResponseWriter, r *http.Request, e error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)

	a.logger.Error(e.Error(), "method", method, "url", url)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Custom error wrapper for client errors
func (a *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
