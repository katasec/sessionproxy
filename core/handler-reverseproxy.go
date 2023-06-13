package core

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (s *Server) reverseProxyHandler(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("In the reverse proxy handler func")

		// Get the claims from auth cookie
		var originHost string
		claims := getCookieClaims(r)
		if claims.Authenticated {
			originHost = "https://www.qnet.net"
			message := fmt.Sprintf("User is authenticated, setting originHost to %s", originHost)
			log.Println(message)

		} else {
			originHost = "https://www.google.com/"
			message := fmt.Sprintf("User is authenticated, setting originHost to %s", originHost)
			log.Println(message)
		}

		reverseProxy := s.getReverseProxy(originHost)

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

func httpLog(r *httputil.ProxyRequest, origin *url.URL) {

	x := fmt.Sprintf("%s://%s%s", origin.Scheme, origin.Host, r.In.URL.String())
	log.Printf("Fetching %s\n", x)

	// Log request headers
	// for k, v := range r.Out.Header {
	// 	log.Printf("Header: %s: %s\n", k, v)
	// }

}
