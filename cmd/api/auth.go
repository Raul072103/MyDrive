package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	// read the auth header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
		return
	}

	// parse it -> get the base64
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Basic" {
		app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
		return
	}

	// decode it
	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		app.unauthorizedBasicErrorResponse(w, r, err)
		return
	}

	// TODO(auth): add retrieval of the password and username
	username := ""
	pass := ""

	// check the credentials
	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 || credentials[0] != username || credentials[1] != pass {
		app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"))
		return
	}
}

func (app *application) createTokenHandler(w http.ResponseWriter, r *http.Request) {

}
