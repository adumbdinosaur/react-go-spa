package main

import (
	"log"
	"net/http"

	v1 "github.com/adumbdinosaur/react-go-spa/server/pkg/api/v1"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Printf("Server started")

	s := v1.Server{}
	router := mux.NewRouter()

	router.Use(cors.AllowAll().Handler)

	r := v1.HandlerWithOptions(s, v1.GorillaServerOptions{
		BaseURL:    "/api/v1",
		BaseRouter: router,
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
