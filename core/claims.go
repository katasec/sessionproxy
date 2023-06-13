package core

import (
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
)

type Claims struct {
	Authenticated bool
	Name          string `json:"name"`
	UserName      string `json:"preferred_username"`
}

func extractTokenClaims(idToken *oidc.IDToken) (claims *Claims, err error) {

	if err := idToken.Claims(&claims); err != nil {
		return &Claims{}, err
	}

	claims.Authenticated = true
	return claims, nil
}

func getCookieClaims(w http.ResponseWriter, r *http.Request) (claims *Claims) {

	var rawIDToken string

	cookies := r.Cookies()

	// Get rawIDToken cookie
	for _, cookie := range cookies {
		if cookie.Name == "rawIDToken" {
			rawIDToken = cookie.Value
		}
	}

	if rawIDToken != "" {
		log.Println("rawIDToken was found: ", rawIDToken)
		idToken, err := auth.Verifier.Verify(ctx, rawIDToken)
		if err != nil {
			log.Println("Failed to verify ID Token: ", idToken)
			log.Println("Error: ", err)
			return &Claims{Authenticated: false}
		}

		claims, err := extractTokenClaims(idToken)
		if err != nil {
			claims.Authenticated = false
		} else {
			claims.Authenticated = true
		}

		return claims
	} else {
		log.Println("rawIDToken was not found: ", rawIDToken)
		return &Claims{Authenticated: false}
	}

}
