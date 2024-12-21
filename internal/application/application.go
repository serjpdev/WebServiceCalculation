package application

import (
	"github.com/serjpdev/WebServiceCalculation/internal/application/logsystem"
	"log/slog"
	"net/http"
	"os"
)

type Config struct {
	Addr string
	Mode string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (a *Application) RunServer() error {
	slog.Info("Start webserver on " + a.config.Addr)
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", logsystem.LogRequestfunc(CalcHandler))
	mux.HandleFunc("/", NotFoundReturn422)
	panicHandler := PanicMiddleware(mux)
	err := http.ListenAndServe(":"+a.config.Addr, panicHandler)

	return err
}
