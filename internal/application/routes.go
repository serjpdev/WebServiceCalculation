package application

import (
	"encoding/json"
	"fmt"
	"github.com/poserj/calc_project/pkg/calculation"
	"io"
	"log/slog"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				bytedata, _ := io.ReadAll(r.Body)
				reqBodyString := string(bytedata)
				slog.Error("start", "method", r.Method, "path", r.URL.Path, "body", reqBodyString)
				http.Error(w, ErrUnknownErrorStr, http.StatusInternalServerError)

			}
		}()
		next.ServeHTTP(w, r)
	})
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrInvalidQueryStr, http.StatusUnprocessableEntity)
		return
	}
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		byteData, _ := io.ReadAll(r.Body)
		reqBodyString := string(byteData)
		slog.Error(err.Error(), "Cant decode request ", reqBodyString)
		http.Error(w, ErrInvalidQueryStr, http.StatusUnprocessableEntity)
		return
	}
	slog.Info("Parse from http request", "expression", request.Expression)
	result, err := calculation.Calc(request.Expression)

	if err != nil {
		slog.Error(err.Error(), "calculation.Calc(request.Expression), it is ", request.Expression)
		http.Error(w, ErrInvalidQueryStr, http.StatusUnprocessableEntity)

	} else {
		slog.Info("Result is", request.Expression, result)
		fmt.Fprintf(w, `{"result": "%f"}`, result)
	}
}
