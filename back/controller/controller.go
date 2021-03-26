package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ricardope3/academy-go-q12021/back/models"
	usecases "github.com/ricardope3/academy-go-q12021/back/usecases"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helloo World")
}

func Pokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		fmt.Print("Url Param 'id' is not given")
	}
	var err error
	requested_id := -1
	if len(ids) > 0 {
		requested_id_str := ids[0]
		requested_id, err = strconv.Atoi(requested_id_str)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	pokemones, errCode := usecases.GetPokemon(requested_id)

	w.WriteHeader(errCode)
	for _, poke := range pokemones {
		json.NewEncoder(w).Encode(poke)
	}

}

func Todos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	todoArray := make([]models.Todo, 0)
	err = json.Unmarshal(bodyBytes, &todoArray)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	errCode := usecases.SaveCSV(todoArray)
	w.WriteHeader(errCode)
	if errCode < 300 {
		json.NewEncoder(w).Encode(todoArray)
	}

}
