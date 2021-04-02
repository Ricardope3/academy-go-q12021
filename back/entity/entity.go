package entity

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ricardope3/academy-go-q12021/back/models"
)

type Entity struct {
}

func New() *Entity {
	return &Entity{}
}

func (e *Entity) GetPokemonFromCSV(requestedId int) ([]models.Pokemon, int) {

	res := make([]models.Pokemon, 0)

	csvFile, err := os.Open("./pokemons.csv")
	if err != nil {
		fmt.Println(err)
		return nil, http.StatusInternalServerError
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, http.StatusInternalServerError
	}

	var found = false
	for _, line := range csvLines {
		idStr := line[0]
		name := line[1]
		id, err := strconv.Atoi(idStr)
		if requestedId > -1 && id != requestedId {
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

func (e *Entity) SaveCSV(todoArray []models.Todo) int {
	csvFile, err := os.Create("./data.csv")
	if err != nil {
		return http.StatusInternalServerError
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
			return http.StatusInternalServerError
		}
	}

	writer.Flush()

	return 202
}

func (e *Entity) GetAllPokemonsFromCSV() ([]models.Pokemon, error) {

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
