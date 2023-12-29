package memory

import (
    "clean-architecture-demo/pkg/entities"
    "errors"
)

type UserMemoryRepository struct {
    users map[int]entities.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
    return &UserMemoryRepository{users: make(map[int]entities.User)}
}

func (repo *UserMemoryRepository) Store(user entities.User) error {
    repo.users[user.ID] = user
    return nil
}

func (repo *UserMemoryRepository) FindByID(id int) (entities.User, error) {
    if user, ok := repo.users[id]; ok {
        return user, nil
    }
    return entities.User{}, errors.New("User not found")
}
