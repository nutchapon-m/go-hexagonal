package ports

import "go-hexagonal/internal/core/domains"

type UserRepository interface {
	Create(user domains.User) error
	GetAll() ([]domains.User, error)
	GetByID(id int) (*domains.User, error)
}
