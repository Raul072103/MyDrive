package main

import (
	"MyDrive/internal/repo"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type SignedInUser struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
}

// loginHandler godoc
//
//	@Summary		Signs in a user
//	@Description	Signs in a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string			true	"Basic authentication (Basic <base64-encoded-credentials>)"
//	@Success		200				{object}	SignedInUser	"User successfully  logged in"
//	@Failure		400				{object}	error			"Invalid request"
//	@Failure		401				{object}	error			"Unauthorized"
//	@Failure		500				{object}	error			"Internal server error"
//	@Router			/auth/login [post]
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

	ctx := r.Context()
	credentials := strings.Split(string(decoded), ":")
	if len(credentials) != 2 {
		app.badRequestResponse(w, r, fmt.Errorf("invalid request"))
		return
	}

	email, pass := credentials[0], credentials[1]

	user, err := app.repo.Users.GetByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrNotFound):
			app.unauthorizedBasicErrorResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	// compare password
	passwordMatch, err := user.Password.Compare(pass)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if !passwordMatch {
		app.unauthorizedBasicErrorResponse(w, r, err)
		return
	}

	// generate the token -> add claims
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(app.config.auth.token.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.token.iss,
		"aud": app.config.auth.token.iss,
	}

	token, err := app.authenticator.GenerateToken(claims)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	signedInUser := SignedInUser{
		Email:     user.Email,
		Username:  user.Username,
		Token:     token,
		TokenType: "bearer",
		ExpiresIn: app.config.auth.token.exp.Milliseconds(),
	}

	if err := jsonResponse(w, http.StatusOK, signedInUser); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) createTokenHandler(w http.ResponseWriter, r *http.Request) {

}
