package controller

import (
	"encoding/json"
	"errors"
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

type safeCounter struct {
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

func WorkerFlags(r *http.Request) (typeStr string, items int, items_per_worker int, err error) {
	type_arr, ok := r.URL.Query()["type"]
	if !ok || len(type_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'type' is not given")
	}
	if len(type_arr) > 0 {
		typeStr = type_arr[0]
		if typeStr != "odd" && typeStr != "even" {
			return "", 0, 0, errors.New("Only support 'odd' or 'even' types")
		}
	}

	items_arr, ok := r.URL.Query()["items"]
	if !ok || len(items_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'items' is not given")
	}
	if len(type_arr) > 0 {
		items_str := items_arr[0]
		items, err = strconv.Atoi(items_str)
		if err != nil {
			return "", 0, 0, errors.New("Items must be an int")
		}
	}

	items_per_worker_arr, ok := r.URL.Query()["items_per_worker"]
	if !ok || len(items_per_worker_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'items_per_worker' is not given")
	}
	err = nil
	if len(items_per_worker_arr) > 0 {
		items_per_worker_str := items_per_worker_arr[0]
		items_per_worker, err = strconv.Atoi(items_per_worker_str)
		if err != nil {
			return "", 0, 0, errors.New("items_per_worker must be an int")
		}
	}
	return typeStr, items, items_per_worker, nil
}

func (c *Controller) Workers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var done = safeCounter{&sync.Mutex{}, 0}

	typeStr, items, maxItemsPerWorker, err := WorkerFlags(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	allPokemonsSlice, err := c.entity.GetAllPokemonsFromCSV()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	values := make(chan models.Pokemon)
	shutdown := make(chan struct{})
	poolSize := items / maxItemsPerWorker
	steps := len(allPokemonsSlice) / poolSize
	var wg sync.WaitGroup
	wg.Add(poolSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < poolSize; i++ {
		go func(ii int) {
			rand.Seed(time.Now().UnixNano())
			starting_index := steps * ii
			ending_index := (steps + steps*ii) - 1
			index := starting_index
			items_of_worker := 0
			for {
				if items_of_worker >= maxItemsPerWorker {

					done.mutex.Lock()
					done.number += 1
					done.mutex.Unlock()
					wg.Done()
					return
				}
				if index >= ending_index {
					wg.Done()
					return
				}
				select {

				case values <- allPokemonsSlice[index]:
					items_of_worker++
					index++

				case <-shutdown:
					wg.Done()
					return
				}

			}
		}(i)
	}
	validPokemons := make([]models.Pokemon, 0)
	for {
		poke := <-values
		if poke.Id%2 == 0 && typeStr != "odd" {
			validPokemons = append(validPokemons, poke)
		} else if poke.Id%2 != 0 && typeStr != "even" {
			validPokemons = append(validPokemons, poke)
		}
		if len(validPokemons) >= items || done.number >= poolSize-1 || len(validPokemons) >= len(allPokemonsSlice) {
			break
		}
	}
	close(shutdown)
	wg.Wait()
	json.NewEncoder(w).Encode(validPokemons)

}
