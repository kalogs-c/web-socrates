package user

import (
	"context"
	"time"

	"github.com/kalogs-c/web-socrates/util"
)

type service struct {
	r       Repository
	timeout time.Duration
}

func NewService(r Repository) Service {
	return &service{r, time.Duration(5) * time.Second}
}

func (s *service) CreateUser(c context.Context, req *UserRequest) (*UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hash, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = hash
	response, err := s.r.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
