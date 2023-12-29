package router

import (
    "clean-architecture-demo/pkg/interfaces/controllers"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type Router struct {
    mux *mux.Router
}

func NewRouter() *Router {
    return &Router{mux: mux.NewRouter()}
}

func (r *Router) SetupRoutes(userController *controllers.UserController) {
    r.mux.HandleFunc("/user", userController.CreateUser).Methods("POST")
    r.mux.HandleFunc("/user", userController.GetUser).Methods("GET")
}

func (r *Router) Start(address string) {
    log.Fatal(http.ListenAndServe(address, r.mux))
}
