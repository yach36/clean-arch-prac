package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/usecase"
	"github.com/yach36/clean-arch-prac/utils/cerrors"
)

type IUserController interface {
	GetUserListHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	PostUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	usecase usecase.IUserUsecase
}

var _ IUserController = (*userController)(nil)

func NewUserController(u usecase.IUserUsecase) *userController {
	return &userController{
		usecase: u,
	}
}

func (c *userController) GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	users, err := c.usecase.GetAllUsers(r.Context())
	if err != nil {
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (c *userController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = cerrors.BadRequest.Wrap(err, "invalid path param")
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}

	user, err := c.usecase.GetUser(r.Context(), id)
	if err != nil {
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *userController) PostUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		err = cerrors.BadRequest.Wrap(err, "invalid body param")
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}

	if err := c.usecase.RegisterUser(r.Context(), user); err != nil {
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewMessage(http.StatusOK, "success"))
}

func (c *userController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = cerrors.BadRequest.Wrap(err, "invalid path param")
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}

	if err := c.usecase.DeleteUser(r.Context(), id); err != nil {
		errorHandler(w, r, cerrors.StatusCode(err), err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewMessage(http.StatusOK, "success"))
}
