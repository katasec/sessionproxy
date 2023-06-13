package core

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func (s *Server) callBack(next ...http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("In the callback handlerfunc")

		// Get "code" query string
		code := r.URL.Query().Get("code")
		log.Println("The code was:", code[0:10])

		// Get the auth session from the request
		oauth2Token, err := auth.Config.Exchange(ctx, code, oauth2.SetAuthURLParam(redirectUriParam, callbackUrl))
		if err != nil {
			message := "Could not get authToken from 'code':" + err.Error()
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
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

		// Create session cookie
		cookie := http.Cookie{
			Name:    "rawIDToken",
			Value:   rawIDToken,
			Expires: oauth2Token.Expiry,
			Path:    "/",
		}
		http.SetCookie(w, &cookie)

		// Redirect to status page
		http.Redirect(w, r, statusUrl, http.StatusFound)

		// Process next middleware
		if (next != nil) && (len(next) > 0) {
			h := next[0]
			h.ServeHTTP(w, r)
		}
	}
}
