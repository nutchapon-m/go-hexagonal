package domains

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
}

type UserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
