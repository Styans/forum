package models

import "time"

type User struct {
	Id        int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	HashedPW  string    `json:"hashed_pw"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserService interface {
	CreateUser(user *CreateUserDTO) error
}
type UserRepo interface {
	CreateUser(user *User) error
}
