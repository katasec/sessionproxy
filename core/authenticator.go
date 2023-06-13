package core

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/url"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider  *oidc.Provider
	Verifier  *oidc.IDTokenVerifier
	Config    oauth2.Config
	Ctx       context.Context
	LogoutUrl string
}

func NewAuthenticator(redirectUrl ...string) (*Authenticator, error) {

	// Create Provider from Azure AD + Tenant ID
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://login.microsoftonline.com/"+azureTenantId+"/v2.0")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	// Create Verifier from Provider
	verifier := provider.Verifier(&oidc.Config{ClientID: azureClientId})

	// Init Logout Url
	logoutUrl, err := url.Parse("https://login.microsoftonline.com/" + azureTenantId + "/oauth2/logout?client_id=" + azureClientId)
	if err != nil {
		log.Printf("failed to parse logout url: %v", err)
		return nil, err
	}

	// Create Config from Provider
	if redirectUrl == nil {
		redirectUrl = []string{callbackUrl}
	}
	conf := oauth2.Config{
		ClientID:     azureClientId,
		ClientSecret: azureClientSecret,
		RedirectURL:  redirectUrl[0],
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider:  provider,
		Config:    conf,
		Ctx:       ctx,
		Verifier:  verifier,
		LogoutUrl: logoutUrl.String(),
	}, nil
}

// generateRandomState Used by the loginHandler to generate a random state
// whilst redirecting to Azure AD. The state is used to prevent CSRF attacks.
func generateRandomState() string {

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)

	return base64.StdEncoding.EncodeToString(b)
}
