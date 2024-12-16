package main

import (
	"github.com/poserj/calc_project/internal/application"
)

func main() {
	app := application.New()
	err := app.RunServer()
	if err != nil {
		panic(err)
	}
}
