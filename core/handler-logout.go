package core

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)

func (s *Server) logoutHandler(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("In the logout handler func")

		//Delete Cookie
		deleteCookie(w, "auth-session")
		deleteCookie(w, "rawIDToken")

		// Redirect to Azure AD logout
		fmt.Println("Redirecitng to: ", azureLogoutUrl)
		http.Redirect(w, r, azureLogoutUrl, http.StatusTemporaryRedirect)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}

	}
}

func deleteCookie(w http.ResponseWriter, name string) {
	cookie := http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Now().Add(-1 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}
