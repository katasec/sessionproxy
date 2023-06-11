package core

import (
	"fmt"
	"net/http"
)

func (s *Server) helloHandler(next ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello World")

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}
