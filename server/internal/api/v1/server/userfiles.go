package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func (s Server) GetUserFiles(w http.ResponseWriter, r *http.Request) {
	username, err := s.Authenticator.Authenticate(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		fmt.Println(err.Error())
		return
	}

	userDir := filepath.Join("static", "files", username)

	files, err := os.ReadDir(userDir)
	if err != nil {
		if os.IsNotExist(err) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string][]string{
				"files": {},
			})
			return
		}
		http.Error(w, `{"error":"failed to read user files"}`, http.StatusInternalServerError)
		return
	}

	fileNames := []string{}
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{
		"files": fileNames,
	})
}

func (s Server) OptionsUserFiles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
