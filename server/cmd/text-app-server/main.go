package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	api "github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/openapi"
	"github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/server"
	"github.com/adumbdinosaur/react-go-spa/server/internal/auth"
)

func main() {
	log.Println("Starting server...")

	authService := auth.NewAuthService()
	apiServer := server.New(authService)
	router := setupRouter(apiServer)

	startServer(router)
}

func setupRouter(apiServer *server.Server) *mux.Router {
	router := mux.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	}).Handler
	router.Use(corsMiddleware)

	api.HandlerWithOptions(*apiServer, api.GorillaServerOptions{
		BaseURL:    "/api/v1",
		BaseRouter: router,
	})

	return router
}

func startServer(handler http.Handler) {
	const port = ":8080"
	log.Printf("Server is running on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
