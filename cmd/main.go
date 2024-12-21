package main

import (
	"github.com/serjpdev/WebServiceCalculation/internal/application"
	"log/slog"
)

func main() {
	app := application.New()
	err := app.RunServer()
	if err != nil {
		slog.Error("Critical error", err)
		panic(err)
	}
}
