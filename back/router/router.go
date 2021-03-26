package router

import (
	"fmt"
	"net/http"

	controller "github.com/ricardope3/academy-go-q12021/back/controller"
)

func Start() {
	http.HandleFunc("/", controller.Root)
	http.HandleFunc("/pokemons", controller.Pokemons)
	http.HandleFunc("/getTodos", controller.Todos)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe("localhost:8000", nil)
}
