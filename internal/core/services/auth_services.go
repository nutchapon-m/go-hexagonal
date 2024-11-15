package services

import (
	"go-hexagonal/internal/core/domains"
	"go-hexagonal/internal/core/ports"
	"go-hexagonal/internal/pkgs/errs"
	"go-hexagonal/internal/pkgs/logs"
)

type authServices struct {
	userRepo ports.UserRepository
}

func NewAuthServices(userRepo ports.UserRepository) ports.AuthService {
	return authServices{userRepo: userRepo}
}

func (s authServices) CreateUser(user domains.UserRequest) error {
	if user.Email != "" && !isValidEmail(user.Email) {
		logs.Error("Invalid email format.")
		return errs.NewError("Invalid email format.")
	}

	userDB := domains.User{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	if err := s.userRepo.Create(userDB); err != nil {
		logs.Error(err)
		return errs.NewError(err.Error())
	}

	return nil
}
