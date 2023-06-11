package user

import (
	"context"

	"github.com/kalogs-c/web-socrates/internal/sqlc"
)

type repository struct {
	q *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) Repository {
	return &repository{q}
}

func (r *repository) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	u, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     req.Name,
		Lastname: req.Lastname,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	response := &UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Lastname: u.Lastname,
		Username: u.Username,
		Email:    u.Email,
	}
	return response, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	u, err := r.q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	response := &UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Lastname: u.Lastname,
		Username: u.Username,
		Email:    u.Email,
	}

	return response, nil
}
