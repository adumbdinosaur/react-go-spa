package server

import (
	"github.com/adumbdinosaur/react-go-spa/server/internal/auth"
)

type Server struct {
	Authenticator *auth.AuthService
}

func New(authenticator *auth.AuthService) *Server {
	return &Server{
		Authenticator: authenticator,
	}
}
