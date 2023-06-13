package core

import (
	"context"
	"log"
	"net/url"
	"os"
)

var (
	azureTenantId     string
	azureClientId     string
	azureClientSecret string

	auth  *Authenticator
	ctx   context.Context
	state string

	logoutUrl        string
	callbackUrl      string
	statusUrl        string
	redirectUriParam string
)

func init() {

	var err error

	// Initialize Context
	ctx = context.Background()

	// Init Azure Tenant Id
	azureTenantId = os.Getenv("AZURE_TENANT_ID")
	if azureTenantId == "" {
		log.Println("AZURE_TENANT_ID not set")
		os.Exit(1)
	}

	// Init Azure Client Id
	azureClientId = os.Getenv("AZURE_CLIENT_ID")
	if azureClientId == "" {
		log.Println("AZURE_CLIENT_ID not set")
		os.Exit(1)
	}

	// Init Azure Client Secret
	azureClientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	if azureClientSecret == "" {
		log.Println("AZURE_CLIENT_SECRET not set")
		os.Exit(1)
	}

	// Initialize Authenticator
	auth, err = NewAuthenticator()
	if err != nil {
		log.Println("Error initalizing authenticator:", err.Error())
		os.Exit(1)
	}

	// Initialize Random State
	state = generateRandomState()

	// Init Logout Url
	mylogoutUrl, err := url.Parse("https://login.microsoftonline.com/" + azureTenantId + "/oauth2/logout?client_id=" + azureClientId)
	if err != nil {
		log.Printf("failed to parse logout url: %v", err)
		os.Exit(1)
	}
	logoutUrl = mylogoutUrl.String()

	// Init Pathfinder Status Url
	statusUrl = "/.pathfinder/status"

	// Callback url
	callbackUrl = "http://localhost:5000/.pathfinder/callback"

	// Redirect Uri Param Name
	redirectUriParam = "redirect_uri"
}
