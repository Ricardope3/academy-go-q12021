package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ricardope3/academy-go-q12021/back/models"
)

type Entity interface {
	GetPokemonFromCSV(requestedId int) ([]models.Pokemon, int)
	SaveCSV(todoArray []models.Todo) int
	GetAllPokemonsFromCSV() ([]models.Pokemon, error)
}

type UseCase interface {
	WorkerFlags(r *http.Request) (string, int, int, error)
}

// Controller struct
type Controller struct {
	entity  Entity
	useCase UseCase
}

// New returns a controller
func New(
	e Entity,
	u UseCase,
) *Controller {
	return &Controller{e, u}
}

type myStruct struct {
	mutex  *sync.Mutex
	number int
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

func (c *Controller) Workers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var done = myStruct{&sync.Mutex{}, 0}

	type_str, items, max_items_per_worker, err := c.useCase.WorkerFlags(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	allPokemonsSlice, err := c.entity.GetAllPokemonsFromCSV()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}

	values := make(chan models.Pokemon)
	shutdown := make(chan struct{})
	poolSize := len(allPokemonsSlice) / 3
	steps := len(allPokemonsSlice) / poolSize
	var wg sync.WaitGroup
	wg.Add(poolSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < poolSize; i++ {
		go func(ii int) {
			rand.Seed(time.Now().UnixNano())
			starting_index := steps * ii
			ending_index := (steps + steps*ii) - 1
			r := rand.Intn(ending_index-starting_index+1) + starting_index
			items_of_worker := 0
			for {
				if items_of_worker >= max_items_per_worker {

					done.mutex.Lock()
					done.number += 1
					done.mutex.Unlock()
					wg.Done()
					return
				}
				select {

				case values <- allPokemonsSlice[r]:
					items_of_worker++

				case <-shutdown:
					wg.Done()
					return
				}

			}
		}(i)
	}
	validPokemons := make([]models.Pokemon, 0)
	numberOfValidPokemons := 0
	for {
		poke := <-values
		if poke.Id%2 == 0 && type_str != "odd" {
			validPokemons = append(validPokemons, poke)
			numberOfValidPokemons++
		} else if poke.Id%2 != 0 && type_str != "even" {
			validPokemons = append(validPokemons, poke)
			numberOfValidPokemons++
		}
		if numberOfValidPokemons >= items || done.number >= poolSize || numberOfValidPokemons >= len(allPokemonsSlice) {
			break
		}
	}
	close(shutdown)
	wg.Wait()
	json.NewEncoder(w).Encode(validPokemons)

}
