package repositories

import (
	"go-hexagonal/internal/core/domains"
	"go-hexagonal/internal/core/ports"

	"gorm.io/gorm"
)

type userRepositories struct {
	db *gorm.DB
}

func NewUserRepositories(db *gorm.DB) ports.UserRepository {

	db.AutoMigrate(&domains.User{})

	return userRepositories{db: db}
}

func (r userRepositories) Create(user domains.User) error {
	return r.db.Create(user).Error
}

func (r userRepositories) GetAll() ([]domains.User, error) {
	panic("Unimplement")
}

func (r userRepositories) GetByID(id int) (*domains.User, error) {
	panic("Unimplement")
}
