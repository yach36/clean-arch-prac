package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yach36/clean-arch-prac/delivery/http/controller"
)

func NewRouter(uc controller.IUserController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", controller.HealthCheckHandler).Methods(http.MethodGet)

	r.HandleFunc("/users", uc.GetUserListHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", uc.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", uc.PostUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", uc.DeleteUserHandler).Methods(http.MethodDelete)

	return r
}
