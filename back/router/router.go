package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	Pokemons(w http.ResponseWriter, r *http.Request)
	Todos(w http.ResponseWriter, r *http.Request)
}

func Start(controller Controller) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/pokemons", controller.Pokemons)
	r.HandleFunc("/getTodos", controller.Todos)
	fmt.Println("Listening on port 8000")
	return r
}
