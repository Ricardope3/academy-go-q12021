package usecases

import (
	"encoding/csv"
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
