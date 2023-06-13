package core

import (
	"log"
	"net/http"
)

func (s *Server) loginHandler(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("In the login handler func")

		// Redirect to Azure AD
		authenticator, err := NewAuthenticator()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Redirecting to: %v", authenticator.Config.AuthCodeURL(state))

		http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}

	}
}
