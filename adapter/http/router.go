package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yach36/clean-arch-prac/adapter/http/controller"
	"github.com/yach36/clean-arch-prac/infra/postgres"
	"github.com/yach36/clean-arch-prac/usecase"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", controller.HealthCheckHandler).Methods(http.MethodGet)

	dbConn := postgres.NewPostgresConnector()
	userRepository := postgres.NewUserRepository(dbConn.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	r.HandleFunc("/users", userController.GetUserListHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userController.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", userController.PostUserHandler).Methods(http.MethodPost)

	return r
}
