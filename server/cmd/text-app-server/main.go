package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	api "github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/openapi"
	"github.com/adumbdinosaur/react-go-spa/server/internal/api/v1/server"
	"github.com/adumbdinosaur/react-go-spa/server/internal/auth"
	"github.com/adumbdinosaur/react-go-spa/server/internal/middleware"
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

	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.SessionMiddleware())

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
