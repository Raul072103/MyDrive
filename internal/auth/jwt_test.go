package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

var jwtAuthenticator *JWTAuthenticator

const Secret = "secret"
const iss = "TES_ISS"
const testUserID = 1

var testExp = time.Now().Add(time.Hour * 24).Unix()
var TestClaims = jwt.MapClaims{
	"sub": testUserID,
	"exp": testExp,
	"iat": time.Now().Unix(),
	"nbf": time.Now().Unix(),
	"iss": iss,
	"aud": iss,
}

func init() {
	jwtAuthenticator = NewJWTAuthenticator(Secret, iss, iss)
}

func TestJWTAuthenticator_TokenCreation(t *testing.T) {
	// Act: Create a token
	token, err := jwtAuthenticator.GenerateToken(TestClaims)

	// Assert: Verify the token creation
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token == "" {
		t.Fatalf("expected a non-empty token")
	}
}

func TestJWTAuthenticator_TokenValidation_ValidToken(t *testing.T) {
	validToken, err := jwtAuthenticator.GenerateToken(TestClaims)
	if err != nil {
		t.Fatalf("failed generating jwt token")
	}

	_, err = jwtAuthenticator.ValidateToken(validToken)

	if err != nil {
		t.Fatalf("expected token to be valid, but got invalid")
	}
}

func TestJWTAuthenticator_TokenValidation_InvalidToken(t *testing.T) {
	invalidToken := "invalid.jwt.token"

	_, err := jwtAuthenticator.ValidateToken(invalidToken)

	if err == nil {
		t.Fatalf("expected token to be invalid")
	}
}

func TestJWTAuthenticator_TokenValidation_TokenWithMissingClaims(t *testing.T) {
	validToken, err := jwtAuthenticator.GenerateToken(jwt.MapClaims{
		"sub": testUserID,
		"exp": testExp,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	if err != nil {
		t.Fatalf("failed generating jwt token")
	}

	_, err = jwtAuthenticator.ValidateToken(validToken)

	if err == nil {
		t.Fatalf("expected token to be invalid, but got valid")
	}
}
