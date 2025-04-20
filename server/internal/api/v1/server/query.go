package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	api "github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/openapi"
	fuzzy "github.com/adumbdinosaur/react-go-spa/server/internal/search"
)

func (s Server) PostQuery(w http.ResponseWriter, r *http.Request) {
	var request api.QueryPostRequest

	username, err := s.Authenticator.Authenticate(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		fmt.Println(err.Error())
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error":"failed to parse request"}`, http.StatusBadRequest)
		return
	}

	if *request.Query == "" || *request.FileName == "" {
		http.Error(w, `{"error":"query and fileName are required"}`, http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("static", "files", username, filepath.Base(*request.FileName))
	fmt.Println("File path:", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, `{"error":"file not found"}`, http.StatusNotFound)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, `{"error":"failed to read file"}`, http.StatusInternalServerError)
		return
	}

	results := fuzzy.FuzzySearch(string(content), *request.Query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"results": results,
	})
}

func (s Server) OptionsQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OptionsQuery called")
}
