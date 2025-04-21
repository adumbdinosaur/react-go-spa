package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	api "github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/openapi"
	"github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/server"
	"github.com/adumbdinosaur/react-go-spa/server/internal/auth"
	"github.com/adumbdinosaur/react-go-spa/server/internal/middleware"
)

func main() {
	// Initialize zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	log.Info().Msg("Starting the server...")

	authService := auth.NewAuthService()
	apiServer := server.New(authService, &log.Logger)

	router := setupRouter(apiServer)

	startServer(router)
}

func setupRouter(apiServer *server.Server) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.SessionMiddleware())
	router.Use(middleware.LoggingMiddleware())

	apiHandler := api.HandlerWithOptions(apiServer, api.GorillaServerOptions{
		BaseRouter: router,
		BaseURL:    "/api/v1",
	})
	router.PathPrefix("/api/v1").Handler(apiHandler)

	return router
}

func startServer(handler http.Handler) {
	port := "8080"
	log.Info().Str("port", port).Msg("Server is running")
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
