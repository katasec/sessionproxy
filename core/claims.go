package core

import oidc "github.com/coreos/go-oidc"

type Claims struct {
	Name     string `json:"name"`
	UserName string `json:"preferred_username"`
}

func extractClaims(idToken *oidc.IDToken) (claims Claims, err error) {

	if err := idToken.Claims(&claims); err != nil {
		return Claims{}, err
	}

	return claims, nil
}
