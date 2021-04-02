package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ricardope3/academy-go-q12021/back/controller"
	"github.com/ricardope3/academy-go-q12021/back/entity"
	app "github.com/ricardope3/academy-go-q12021/back/router"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	entity := entity.New()
	controller := controller.New(entity)
	r := app.Start(controller)
	port := os.Getenv("PORT")
	http.ListenAndServe("localhost:"+port, r)

}
