package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yach36/clean-arch-prac/adapter/http/controller"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", controller.HealthCheckHandler).Methods(http.MethodGet)

	// r.Use(middleware.LoggingMiddleware)

	return r
}
