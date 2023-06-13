package core

import (
	"context"
	"log"
	"net/url"
	"os"

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
	provider, err := oidc.NewProvider(ctx, "https://login.microsoftonline.com/"+os.Getenv("AZURE_TENANT_ID")+"/v2.0")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	// Create Verifier from Provider
	verifier := provider.Verifier(&oidc.Config{ClientID: os.Getenv("AZURE_CLIENT_ID")})

	// Init Logout Url
	logoutUrl, err := url.Parse("https://login.microsoftonline.com/" + os.Getenv("AZURE_TENANT_ID") + "/oauth2/logout?client_id=" + os.Getenv("AZURE_CLIENT_ID"))
	if err != nil {
		log.Printf("failed to parse logout url: %v", err)
		return nil, err
	}

	// Create Config from Provider
	if redirectUrl == nil {
		redirectUrl = []string{"http://localhost:5000/.pathfinder/callback"}
	}
	conf := oauth2.Config{
		ClientID:     os.Getenv("AZURE_CLIENT_ID"),
		ClientSecret: os.Getenv("AZURE_CLIENT_SECRET"),
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
