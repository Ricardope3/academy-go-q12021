package usecases

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ricardope3/academy-go-q12021/back/models"
)

type UseCase struct {
}

// New community UseCase
func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) GetPokemon(requested_id int) ([]models.Pokemon, int) {

	res := make([]models.Pokemon, 0)

	csvFile, err := os.Open("./pokemons.csv")
	if err != nil {
		fmt.Println(err)
		return nil, 500
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, 500
	}

	var found = false
	for _, line := range csvLines {
		id_str := line[0]
		name := line[1]
		id, err := strconv.Atoi(id_str)
		if requested_id > -1 && id != requested_id {
			continue
		}
		found = true
		if err != nil {
			fmt.Println(err)
		}
		poke := models.Pokemon{
			Name: name,
			Id:   id,
		}
		res = append(res, poke)
	}

	if !found {
		fmt.Println("No pokemon found with given ID")
		return nil, http.StatusNotFound
	}
	return res, 202

}

func (u *UseCase) SaveCSV(todoArray []models.Todo) int {
	csvFile, err := os.Create("./data.csv")
	if err != nil {
		return 500
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, todo := range todoArray {
		var row []string
		row = append(row, strconv.Itoa(todo.ID))
		row = append(row, todo.Title)
		row = append(row, strconv.Itoa(todo.Userid))
		row = append(row, strconv.FormatBool(todo.Completed))
		err = writer.Write(row)
		if err != nil {
			return 500
		}
	}

	writer.Flush()

	return 202
}

func (u *UseCase) WorkerFlags(r *http.Request) (string, int, int, error) {
	type_arr, ok := r.URL.Query()["type"]
	var err error
	var type_str = ""
	var items = -1
	var items_per_worker = -1
	if !ok || len(type_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'type' is not given")
	}
	if len(type_arr) > 0 {
		type_str = type_arr[0]
		if type_str != "odd" && type_str != "even" {
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
	return type_str, items, items_per_worker, nil
}

func (u *UseCase) GetAllPokemons() ([]models.Pokemon, error) {

	res := make([]models.Pokemon, 0)

	csvFile, err := os.Open("./pokemons.csv")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Cant Open CSV")
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Cant Read CSV")
	}

	for _, line := range csvLines {
		id_str := line[0]
		name := line[1]
		id, err := strconv.Atoi(id_str)
		if err != nil {
			fmt.Println(err)
		}
		poke := models.Pokemon{
			Name: name,
			Id:   id,
		}
		res = append(res, poke)
	}

	return res, nil

}
