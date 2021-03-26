package main

import (
	"net/http"

	"github.com/ricardope3/academy-go-q12021/back/controller"
	app "github.com/ricardope3/academy-go-q12021/back/router"
	"github.com/ricardope3/academy-go-q12021/back/usecases"
)

func main() {
	useCase := usecases.New()
	controller := controller.New(useCase)
	r := app.Start(controller)

	http.ListenAndServe("localhost:8000", r)

}
