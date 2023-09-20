package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yach36/clean-arch-prac/usecase"
)

type userController struct {
	usecase usecase.IUserUsecase
}

func NewUserController(u usecase.IUserUsecase) *userController {
	return &userController{
		usecase: u,
	}
}

func (c *userController) GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	users, err := c.usecase.GetAllUsers(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}

func (c *userController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	user, err := c.usecase.GetUser(r.Context(), id)
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
