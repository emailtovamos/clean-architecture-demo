package main

import (
    "clean-architecture-demo/pkg/infrastructure/router"
    "clean-architecture-demo/pkg/interfaces/controllers"
    "clean-architecture-demo/pkg/infrastructure/datastore/memory"
    "clean-architecture-demo/pkg/usecases"
)

func main() {
    userRepo := memory.NewUserMemoryRepository()
    userInteractor := usecases.NewUserInteractor(userRepo)
    userController := controllers.NewUserController(*userInteractor)

    router := router.NewRouter()
    router.SetupRoutes(userController)
    router.Start(":8080")
}
