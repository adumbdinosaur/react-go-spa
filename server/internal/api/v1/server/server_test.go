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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostRegister(t *testing.T) {
	authService := auth.NewAuthService()
	server := New(authService)

	// Create a mock HTTP request
	body := api.RegisterPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Call the handler
	server.PostRegister(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, rec.Code)
	var resp api.RegisterPost201Response
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Token)
}

func TestPostLogin(t *testing.T) {
	authService := auth.NewAuthService()
	server := New(authService)

	// Register a user using the server's PostRegister method
	registerBody := api.RegisterPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	registerBodyBytes, _ := json.Marshal(registerBody)
	registerReq := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(registerBodyBytes))
	registerReq.Header.Set("Content-Type", "application/json")
	registerRec := httptest.NewRecorder()
	server.PostRegister(registerRec, registerReq)

	// Assert that registration was successful
	assert.Equal(t, http.StatusCreated, registerRec.Code)

	// Create a mock HTTP request for login
	loginBody := api.LoginPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	loginBodyBytes, _ := json.Marshal(loginBody)
	loginReq := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(loginBodyBytes))
	loginReq.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	loginRec := httptest.NewRecorder()

	// Call the handler
	server.PostLogin(loginRec, loginReq)

	// Assert the response
	assert.Equal(t, http.StatusOK, loginRec.Code)
	var resp api.LoginPost200Response
	err := json.Unmarshal(loginRec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Token)
}

func TestPostQuery(t *testing.T) {
	authService := auth.NewAuthService()
	server := New(authService)

	// Register a user
	registerBody := api.RegisterPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	registerBodyBytes, _ := json.Marshal(registerBody)
	registerReq := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(registerBodyBytes))
	registerReq.Header.Set("Content-Type", "application/json")
	registerRec := httptest.NewRecorder()
	server.PostRegister(registerRec, registerReq)
	require.Equal(t, http.StatusCreated, registerRec.Code)

	// Log in to get a token
	loginBody := api.LoginPostRequest{
		Username: stringPtr("testuser"),
		Password: stringPtr("password123"),
	}
	loginBodyBytes, _ := json.Marshal(loginBody)
	loginReq := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(loginBodyBytes))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	server.PostLogin(loginRec, loginReq)
	require.Equal(t, http.StatusOK, loginRec.Code)

	var loginResp api.LoginPost200Response
	err := json.Unmarshal(loginRec.Body.Bytes(), &loginResp)
	require.NoError(t, err)
	require.NotNil(t, loginResp.Token)

	// Create the required file in the test environment
	username := "testuser"
	filePath := filepath.Join("static", "files", username)
	err = os.MkdirAll(filePath, os.ModePerm)
	require.NoError(t, err)

	testFilePath := filepath.Join(filePath, "test.txt")
	err = os.WriteFile(testFilePath, []byte("This is a test file content"), 0644)
	require.NoError(t, err)

	// Create a mock HTTP request for query
	queryBody := api.QueryPostRequest{
		Query:    stringPtr("test"),
		FileName: stringPtr("test.txt"),
	}
	queryBodyBytes, _ := json.Marshal(queryBody)
	queryReq := httptest.NewRequest(http.MethodPost, "/query", bytes.NewReader(queryBodyBytes))
	queryReq.Header.Set("Content-Type", "application/json")
	queryReq.Header.Set("Authorization", "Bearer "+*loginResp.Token)

	// Create a response recorder
	queryRec := httptest.NewRecorder()

	// Call the handler
	server.PostQuery(queryRec, queryReq)

	// Assert the response
	require.Equal(t, http.StatusOK, queryRec.Code)
	var queryResp api.QueryPost200Response
	err = json.Unmarshal(queryRec.Body.Bytes(), &queryResp)
	require.NoError(t, err)
	require.NotNil(t, queryResp.Results)

	// Clean up the test file
	err = os.RemoveAll("static")
	require.NoError(t, err)
}

func stringPtr(s string) *string {
	return &s
}
