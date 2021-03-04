package app

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type pokemon struct {
	Id   int
	Name string
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helloo World")

}

func pokemons(w http.ResponseWriter, r *http.Request) {

	csvFile, err := os.Open("./pokemons.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		fmt.Print("Url Param 'id' is not given")
	}
	requested_id := -1
	if len(ids) > 0 {
		requested_id_str := ids[0]
		requested_id, err = strconv.Atoi(requested_id_str)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, line := range csvLines {
		id_str := line[0]
		name := line[1]
		id, err := strconv.Atoi(id_str)
		if requested_id != -1 && id != requested_id {
			continue
		}

		if err != nil {
			fmt.Println(err)
		}
		poke := pokemon{
			Name: name,
			Id:   id,
		}
		fmt.Fprint(w, poke.Id, " ", poke.Name, "\n")
	}

}
