package server

import (
	"log"
	"net/http"

	openapi "github.com/adumbdinosaur/react-go-spa/server/pkg/api/v1"
)

func main() {
	log.Printf("Server started")

	DefaultAPIService := openapi.NewDefaultAPIService()
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
