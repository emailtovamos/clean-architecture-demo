package gateways

import (
    "clean-architecture-demo/pkg/entities"
)

type UserRepository interface {
    Store(user entities.User) error
    FindByID(id int) (entities.User, error)
}
