package router

import (
	"github.com/alpgozbasi/image-processing-service/internal/config"
	"github.com/alpgozbasi/image-processing-service/internal/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	r.HandleFunc("/api/v1/convert", handler.ConvertImageHandler(cfg)).Methods(http.MethodPost)

	return r
}
