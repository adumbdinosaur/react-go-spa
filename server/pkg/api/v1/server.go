package v1

import (
	"fmt"
	"net/http"
)

type Server struct{}

func (s Server) PostQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostQuery called")

}

func (s Server) PostUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostUpload called")
}

func (s Server) OptionsQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OptionsQuery called")
}

func (s Server) OptionsUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OptionsUpload called")
}
