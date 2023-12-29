package usecases

import (
    "clean-architecture-demo/pkg/entities"
)

type UserRepository interface {
    Store(user entities.User) error
    FindByID(id int) (entities.User, error)
}

type UserInteractor struct {
    UserRepository UserRepository
}

func NewUserInteractor(repository UserRepository) *UserInteractor {
    return &UserInteractor{UserRepository: repository}
}

func (interactor *UserInteractor) AddUser(user entities.User) error {
    return interactor.UserRepository.Store(user)
}

func (interactor *UserInteractor) GetUser(id int) (entities.User, error) {
    return interactor.UserRepository.FindByID(id)
}
