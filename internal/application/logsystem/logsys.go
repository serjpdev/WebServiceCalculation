package logsystem

import (
	"log/slog"
	"net/http"
)

func LogRequestfunc(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		slog.Info("start", "method", r.Method, "path", r.URL.Path)
		next(w, r)
	}
}
