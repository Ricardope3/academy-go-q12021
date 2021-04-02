package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	Pokemons(w http.ResponseWriter, r *http.Request)
	Todos(w http.ResponseWriter, r *http.Request)
	Workers(w http.ResponseWriter, r *http.Request)
}

func Start(controller Controller) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/pokemons", controller.Pokemons)
	r.HandleFunc("/todos", controller.Todos)
	r.HandleFunc("/workers", controller.Workers)
	fmt.Println("Listening on port 8000")
	return r
}
