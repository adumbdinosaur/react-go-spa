package server

import (
	"net/http"
)

func (s Server) PostLogin(w http.ResponseWriter, r *http.Request) {
	s.Authenticator.Login(w, r)
}

func (s Server) OptionsLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
