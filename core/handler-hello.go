package core

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *Server) helloHandlerFunc(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("In the hello handler func")
		var message string

		claims := getCookieClaims(w, r)

		tmStamp := time.Now().Format("2006-01-02 15:04:05")

		if claims.Authenticated {
			message = fmt.Sprintf("%s, Hello %s (%s)", tmStamp, claims.Name, claims.UserName)
		} else {
			message = fmt.Sprintf("%s Hello World, your anonymous !", tmStamp)
		}

		fmt.Fprintln(w, message)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}

func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("In the hello handler func")
	var message string

	claims := getCookieClaims(w, r)

	tmStamp := time.Now().Format("2006-01-02 15:04:05")

	if claims.Authenticated {
		message = fmt.Sprintf("%s, Hello %s (%s)", tmStamp, claims.Name, claims.UserName)
	} else {
		message = fmt.Sprintf("%s Hello World, your anonymous !", tmStamp)
	}

	fmt.Fprintln(w, message)

}
