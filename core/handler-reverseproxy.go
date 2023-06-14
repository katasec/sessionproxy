package core

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	// authOrigin = "https://www.qnet.net"
	// authProxy  = getReverseProxy(authOrigin)

	// anonOrigin = "https://www.google.com/"
	// anonProxy  = getReverseProxy(anonOrigin)
	authOrigin string
	authProxy  *httputil.ReverseProxy

	anonOrigin string
	anonProxy  *httputil.ReverseProxy
)

func init() {

	getOrginFromEnv()

}

func (s *Server) reverseProxyHandler(next ...http.HandlerFunc) http.HandlerFunc {

	//getOrginFromEnv()

	return func(w http.ResponseWriter, r *http.Request) {

		var reverseProxy *httputil.ReverseProxy

		log.Println("In the reverse proxy handler func")

		// Get the claims from auth cookie
		claims := getCookieClaims(r)
		if claims.Authenticated {
			message := fmt.Sprintf("User is authenticated, using originHost:%s", authOrigin)
			reverseProxy = authProxy
			log.Println(message)

		} else {
			message := fmt.Sprintf("User is anonymous, using originHost:%s", anonOrigin)
			reverseProxy = anonProxy
			log.Println(message)
		}

		//reverseProxy := s.getReverseProxy(originHost)

		// Process the request through the reverse proxy
		reverseProxy.ServeHTTP(w, r)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}

// getReverseProxy Create a reverse proxy forwardwing to the appropriate target
// host for a given service.
func (s *Server) getReverseProxy(originHost string) *httputil.ReverseProxy {

	origin, _ := url.Parse(originHost)

	return &httputil.ReverseProxy{

		Rewrite: func(r *httputil.ProxyRequest) {

			// Set the request host to the target host
			r.SetURL(origin)

			// Log request
			httpLog(r, origin)
		},
	}
}

// getReverseProxy Create a reverse proxy forwardwing to the appropriate target
// host for a given service.
func getReverseProxy(originHost string) *httputil.ReverseProxy {

	origin, _ := url.Parse(originHost)

	return &httputil.ReverseProxy{

		Rewrite: func(r *httputil.ProxyRequest) {

			// Set the request host to the target host
			r.SetURL(origin)

			// Log request
			httpLog(r, origin)
		},
	}
}

func httpLog(r *httputil.ProxyRequest, origin *url.URL) {

	x := fmt.Sprintf("%s://%s%s", origin.Scheme, origin.Host, r.In.URL.String())
	log.Printf("Fetching %s\n", x)

	// Log request headers
	// for k, v := range r.Out.Header {
	// 	log.Printf("Header: %s: %s\n", k, v)
	// }

}

func getOrginFromEnv() {
	// Set origin host for authenticated users
	authOrigin = os.Getenv("SPROXY_AUTH_ORIGIN")
	if authOrigin == "" {
		log.Println("SPROXY_AUTH_ORIGIN not set, exitting...")
		os.Exit(1)
	}
	authProxy = getReverseProxy(authOrigin)

	// Set origin host for anonymous  users
	anonOrigin = os.Getenv("SPROXY_ANON_ORIGIN")
	if anonOrigin == "" {
		log.Println("SPROXY_ANON_ORIGIN not set, exitting...")
		os.Exit(1)
	}
	anonProxy = getReverseProxy(anonOrigin)
}
