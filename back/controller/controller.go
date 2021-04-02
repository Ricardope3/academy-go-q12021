package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ricardope3/academy-go-q12021/back/models"
)

type UseCase interface {
	GetPokemon(requested_id int) ([]models.Pokemon, int)
	SaveCSV(todoArray []models.Todo) int
}

// Controller struct
type Controller struct {
	useCase UseCase
}

// New returns a controller
func New(
	u UseCase,
) *Controller {
	return &Controller{u}
}


func (c *Controller) Pokemons(w http.ResponseWriter, r *http.Request) {
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

	pokemones, errCode := c.useCase.GetPokemon(requested_id)

	w.WriteHeader(errCode)
	for _, poke := range pokemones {
		json.NewEncoder(w).Encode(poke)
	}

}

func (c *Controller) Todos(w http.ResponseWriter, r *http.Request) {
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
	errCode := c.useCase.SaveCSV(todoArray)
	w.WriteHeader(errCode)
	if errCode < 300 {
		json.NewEncoder(w).Encode(todoArray)
	}

}
