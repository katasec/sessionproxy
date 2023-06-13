package core

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) reverseProxyHandler(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("In the logout handler func")

		//Delete Cookie
		deleteCookie(w, "auth-session")
		deleteCookie(w, "rawIDToken")

		// Redirect to Azure AD logout
		fmt.Println("Redirecitng to: ", logoutUrl)
		http.Redirect(w, r, logoutUrl, http.StatusTemporaryRedirect)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}
