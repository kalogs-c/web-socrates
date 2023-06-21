package user

import (
	"context"
	"fmt"

	"github.com/lib/pq"

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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, fmt.Errorf("user already exists")
			}
		}
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

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*sqlc.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
