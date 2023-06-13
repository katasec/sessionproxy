package core

import "net/http"

func nextHandler(w http.ResponseWriter, r *http.Request, next ...http.HandlerFunc) {
	if (next != nil) && (len(next) > 0) {
		h := next[0]
		h.ServeHTTP(w, r)
	}
}
