package user

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/kalogs-c/web-socrates/internal/sqlc"
)

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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string       `json:"token"`
	User      UserResponse `json:"user"`
	ClientIP  string       `json:"client_ip"`
	UserAgent string       `json:"user_agent"`
}

type Repository interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*sqlc.User, error)
}

type Service interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func (ur *UserRequest) Validate() error {
	if ur.Name == "" {
		return fmt.Errorf("name is required")
	}
	if ur.Lastname == "" {
		return fmt.Errorf("lastname is required")
	}
	if ur.Username == "" {
		return fmt.Errorf("username is required")
	}
	if ur.Email == "" {
		return fmt.Errorf("email is required")
	}
	if ur.Password == "" || len(ur.Password) < 5 {
		return fmt.Errorf("password is required")
	}

	if _, err := mail.ParseAddress(ur.Email); err != nil {
		return fmt.Errorf("email is invalid")
	}

	return nil
}
