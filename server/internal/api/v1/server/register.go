package server

import (
	"net/http"
)

func (s Server) PostRegister(w http.ResponseWriter, r *http.Request) {
	s.Authenticator.Register(w, r)
}

func (s Server) OptionsRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
