package main

import (
	"github.com/serjpdev/WebServiceCalculation/internal/application"
)

func main() {
	app := application.New()
	err := app.RunServer()
	if err != nil {
		panic(err)
	}
}
