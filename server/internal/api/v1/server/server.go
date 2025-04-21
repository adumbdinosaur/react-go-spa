package server

import (
	"github.com/adumbdinosaur/react-go-spa/server/internal/auth"
	"github.com/rs/zerolog"
)

type Server struct {
	Authenticator *auth.AuthService
	Logger        *zerolog.Logger
}

func New(authenticator *auth.AuthService, logger *zerolog.Logger) *Server {
	return &Server{
		Authenticator: authenticator,
		Logger:        logger,
	}
}
