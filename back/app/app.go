package app

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/", root)
	http.HandleFunc("/pokemons", pokemons)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe("localhost:8000", nil)
}
