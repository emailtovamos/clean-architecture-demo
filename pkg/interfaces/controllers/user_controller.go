package controllers

import (
    "clean-architecture-demo/pkg/entities"
    "clean-architecture-demo/pkg/usecases"
    "net/http"
    "encoding/json"
    "strconv"
)

type UserController struct {
    Interactor usecases.UserInteractor
}

func NewUserController(interactor usecases.UserInteractor) *UserController {
    return &UserController{Interactor: interactor}
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user entities.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    err := controller.Interactor.AddUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    user, err := controller.Interactor.GetUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}
