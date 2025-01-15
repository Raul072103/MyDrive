package main

import (
	"MyDrive/internal/auth"
	"MyDrive/internal/repo"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T, cfg config, repoBuilder repo.IMockRepositoryBuilder) *application {
	t.Helper()
	logger := zap.Must(zap.NewDevelopment()).Sugar()

	mockRepo := repoBuilder.GetRepository()

	testAuth := auth.NewTestAuthenticator()

	return &application{
		logger:        logger,
		repo:          *mockRepo,
		authenticator: testAuth,
		config:        cfg,
	}
}

func executeRequest(req *http.Request, mux *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected response code %d but got %d", expected, actual)
	}
}
