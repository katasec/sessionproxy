package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (s *Server) callBack(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("In the status handlerfunc")

		// Get "code" query string
		code := r.URL.Query().Get("code")
		log.Println("The code was:", code[0:10])

		// Get the auth session from the request
		oauth2Token, err := auth.Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			log.Println("Could not extract id_token from oauth2 token.")
			os.Exit(1)
		}

		// Parse and verify ID Token payload.
		idToken, err := auth.Verifier.Verify(ctx, rawIDToken)
		if err != nil {
			log.Println("Failed to verify ID Token: ", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
			//return
		}

		// Extract custom claims
		claims, err := extractClaims(idToken)
		if err != nil {
			message := "Failed to extract custom claims: " + err.Error()
			http.Error(w, message+err.Error(), http.StatusInternalServerError)
			log.Println(message + err.Error())
		}

		// Response with hello world
		message := "Hello " + claims.Name + " (" + claims.UserName + ")"
		fmt.Fprintln(w, message)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}
