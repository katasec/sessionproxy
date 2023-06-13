package core

import (
	"log"
	"net/http"
	"os"
)

func (s *Server) callBack(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("In the callback handlerfunc")

		// Get "code" query string
		code := r.URL.Query().Get("code")
		log.Println("The code was:", code[0:10])

		// Get the auth session from the request
		oauth2Token, err := auth.Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			message := "Could not extract id_token from oauth2 token:" + err.Error()
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			os.Exit(1)
		}

		// Parse and verify ID Token payload.
		// idToken, err := auth.Verifier.Verify(ctx, rawIDToken)
		// if err != nil {
		// 	log.Println("Failed to verify ID Token: ", err)
		// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	log.Println(err.Error())
		// 	//return
		// }

		// Create session cookie
		cookie := http.Cookie{
			Name:    "rawIDToken",
			Value:   rawIDToken,
			Expires: oauth2Token.Expiry,
			Path:    "/",
		}
		http.SetCookie(w, &cookie)

		// Redirect to home page
		http.Redirect(w, r, "/", http.StatusFound)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}
