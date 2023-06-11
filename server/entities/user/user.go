package user

import "context"

type UserRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Repository interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}

type Service interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}
