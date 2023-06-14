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

	azureLogoutUrl   string
	callbackUrl      string
	redirectUriParam string

	SPROXY_PORT string

	pfLoginUrl    string
	pfLogoutUrl   string
	pfStatusUrl   string
	pfProxyPath   string
	pfCallbackUrl string
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

	// Init listen port SPROXY_PORT
	SPROXY_PORT = os.Getenv("SPROXY_PORT")
	if SPROXY_PORT == "" {
		log.Println("SPROXY_PORT not set, defaulting to 8080")
		SPROXY_PORT = "8080"
	}

	// Initialize Authenticator
	auth, err = NewAuthenticator()
	if err != nil {
		log.Println("Error initalizing authenticator:", err.Error())
		os.Exit(1)
	}

	// Initialize Random State
	state = generateRandomState()

	// Init Azure Logout Url
	mylogoutUrl, err := url.Parse("https://login.microsoftonline.com/" + azureTenantId + "/oauth2/logout?client_id=" + azureClientId)
	if err != nil {
		log.Printf("failed to parse logout url: %v", err)
		os.Exit(1)
	}
	azureLogoutUrl = mylogoutUrl.String()

	// Init OAuth2 Callback Url
	callbackUrl = "https://portal.qntest.com/ameer/.pathfinder/callback"
	log.Println("Call back url is:", callbackUrl)

	// Init Pathfinder Login Url
	pfLoginUrl = "/ameer/.pathfinder/login"

	// Init Pathfinder Logout Url
	pfLogoutUrl = "/ameer/.pathfinder/logout"

	// Init Pathfinder Status Url
	pfStatusUrl = "/ameer/.pathfinder/status"

	// Init Pathfinder Status Url
	pfCallbackUrl = "/ameer/.pathfinder/callback"

	// Init Proxy Path
	pfProxyPath = "/ameer"

	// Redirect Uri Param Name
	redirectUriParam = "redirect_uri"
}
