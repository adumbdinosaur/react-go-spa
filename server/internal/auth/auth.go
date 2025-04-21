package auth

import (
	"encoding/json"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"

	api "github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/openapi"
	"github.com/adumbdinosaur/react-go-spa/server/internal/middleware"
)

type AuthService struct {
	users map[string]string
	mu    sync.Mutex
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]string),
	}
}

func (a *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req api.RegisterPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if *req.Username == "" || *req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if _, exists := a.users[*req.Username]; exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	a.users[*req.Username] = string(hashedPassword)

	session, err := middleware.SessionStore.Get(r, middleware.SessionName)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	session.Values["username"] = *req.Username
	session.Values["authenticated"] = true

	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req api.LoginPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if *req.Username == "" || *req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	hashedPassword, exists := a.users[*req.Username]
	if !exists {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(*req.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	session, err := middleware.SessionStore.Get(r, middleware.SessionName)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	session.Values["username"] = *req.Username
	session.Values["authenticated"] = true

	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (a *AuthService) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	session, err := middleware.SessionStore.Get(r, middleware.SessionName)
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to clear session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (a *AuthService) Authenticate(r *http.Request) (string, error) {
	session, err := middleware.SessionStore.Get(r, middleware.SessionName)
	if err != nil {
		return "", http.ErrNoCookie
	}

	authenticated, ok := session.Values["authenticated"].(bool)
	if !ok || !authenticated {
		return "", http.ErrNoCookie
	}

	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		return "", http.ErrNoCookie
	}

	return username, nil
}
