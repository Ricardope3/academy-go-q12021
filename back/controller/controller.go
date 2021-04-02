package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ricardope3/academy-go-q12021/back/models"
)

type Entity interface {
	GetPokemonFromCSV(requestedId int) ([]models.Pokemon, int)
	SaveCSV(todoArray []models.Todo) int
}

// Controller struct
type Controller struct {
	entity Entity
}

// New returns a controller
func New(
	u Entity,
) *Controller {
	return &Controller{u}
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helloo World")
}

func (c *Controller) Pokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		fmt.Print("Url Param 'id' is not given")
	}
	var err error
	requestedId := -1
	if len(ids) > 0 {
		requestedIdStr := ids[0]
		requestedId, err = strconv.Atoi(requestedIdStr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	pokemones, errCode := c.entity.GetPokemonFromCSV(requestedId)

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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	todoArray := make([]models.Todo, 0)
	err = json.Unmarshal(bodyBytes, &todoArray)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	errCode := c.entity.SaveCSV(todoArray)
	w.WriteHeader(errCode)
	if errCode < 300 {
		json.NewEncoder(w).Encode(todoArray)
	}

}
