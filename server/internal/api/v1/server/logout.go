package server

import "net/http"

func (s Server) PostLogout(w http.ResponseWriter, r *http.Request) {
	s.Authenticator.LogOut(w, r)
}

func (s Server) OptionsLogout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
