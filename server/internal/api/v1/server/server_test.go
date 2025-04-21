package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	api "github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/openapi"
	"github.com/adumbdinosaur/react-go-spa/server/internal/auth"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostRegister(t *testing.T) {
	authService := auth.NewAuthService()
	logger := newTestLogger(t)
	server := New(authService, &logger)

	body := api.RegisterPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	bodyBytes, err := json.Marshal(body)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.PostRegister(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestPostLogin(t *testing.T) {
	authService := auth.NewAuthService()
	logger := newTestLogger(t)
	server := New(authService, &logger)

	registerBody := api.RegisterPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	registerBodyBytes, err := json.Marshal(registerBody)
	require.NoError(t, err)

	registerReq := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(registerBodyBytes))
	registerReq.Header.Set("Content-Type", "application/json")
	registerRec := httptest.NewRecorder()

	server.PostRegister(registerRec, registerReq)

	require.Equal(t, http.StatusCreated, registerRec.Code)

	loginBody := api.LoginPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	loginBodyBytes, err := json.Marshal(loginBody)
	require.NoError(t, err)

	loginReq := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(loginBodyBytes))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()

	server.PostLogin(loginRec, loginReq)

	assert.Equal(t, http.StatusOK, loginRec.Code)
}

func TestPostQuery(t *testing.T) {
	authService := auth.NewAuthService()
	logger := newTestLogger(t)
	server := New(authService, &logger)

	registerBody := api.RegisterPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	registerBodyBytes, err := json.Marshal(registerBody)
	require.NoError(t, err)

	registerReq := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(registerBodyBytes))
	registerReq.Header.Set("Content-Type", "application/json")
	registerRec := httptest.NewRecorder()

	server.PostRegister(registerRec, registerReq)

	require.Equal(t, http.StatusCreated, registerRec.Code)

	loginBody := api.LoginPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	loginBodyBytes, err := json.Marshal(loginBody)
	require.NoError(t, err)

	loginReq := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(loginBodyBytes))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()

	server.PostLogin(loginRec, loginReq)

	require.Equal(t, http.StatusOK, loginRec.Code)

	username := "testuser"
	filePath := filepath.Join("static", "files", username)
	err = os.MkdirAll(filePath, os.ModePerm)
	require.NoError(t, err)

	testFilePath := filepath.Join(filePath, "test.txt")
	err = os.WriteFile(testFilePath, []byte("This is a test file content"), 0644)
	require.NoError(t, err)

	queryBody := api.QueryPostRequest{
		Query:    stringPtr("test"),
		FileName: stringPtr("test.txt"),
	}
	queryBodyBytes, err := json.Marshal(queryBody)
	require.NoError(t, err)

	queryReq := httptest.NewRequest(http.MethodPost, "/query", bytes.NewReader(queryBodyBytes))
	queryReq.Header.Set("Content-Type", "application/json")
	queryReq.Header.Set("Cookie", loginRec.Header().Get("Set-Cookie"))
	queryRec := httptest.NewRecorder()

	server.PostQuery(queryRec, queryReq)

	require.Equal(t, http.StatusOK, queryRec.Code)
	var queryResp api.QueryPost200Response
	err = json.Unmarshal(queryRec.Body.Bytes(), &queryResp)
	require.NoError(t, err)
	require.NotNil(t, queryResp.Results)

	err = os.RemoveAll("static")
	require.NoError(t, err)
}

func stringPtr(s string) *string {
	return &s
}

func newTestLogger(t *testing.T) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:     os.Stderr,
		NoColor: true,
	}).With().Timestamp().Logger()
}
