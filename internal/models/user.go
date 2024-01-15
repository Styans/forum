package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	HashedPW  string    `json:"hashed_pw"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

type UserService interface {
	CreateUser(user *CreateUserDTO) error
}

type UserRepo interface {
	CreateUser(user *User) error
	
}
