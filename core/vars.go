package core

import (
	"context"
	"fmt"
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

	SPROXY_PORT string
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
	SPROXY_PORT = os.Getenv("SPROXY_PORT")
	if SPROXY_PORT == "" {
		log.Println("SPROXY_PORT not set, defaulting to 8080")
		SPROXY_PORT = "8080"
	}
	callbackUrl = fmt.Sprintf("http://localhost:%s/.pathfinder/callback", SPROXY_PORT)
	log.Println("Call back url is:", callbackUrl)

	// Redirect Uri Param Name
	redirectUriParam = "redirect_uri"
}
