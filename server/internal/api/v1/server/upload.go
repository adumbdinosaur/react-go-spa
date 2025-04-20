package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (s Server) PostUpload(w http.ResponseWriter, r *http.Request) {
	username, err := s.Authenticator.Authenticate(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		fmt.Println(err.Error())
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, `{"error":"failed to parse form"}`, http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, `{"error":"file is required"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := filepath.Base(header.Filename)
	userDir := filepath.Join("static", "files", username)

	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		http.Error(w, `{"error":"failed to create user directory"}`, http.StatusInternalServerError)
		return
	}
	target := filepath.Join(userDir, filename)

	out, err := os.Create(target)
	if err != nil {
		http.Error(w, `{"error":"cannot create file"}`, http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, `{"error":"failed to save file"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "file uploaded",
		"filename": filename,
		"path":     target,
	})
}

func (s Server) OptionsUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OptionsUpload called")
}
