package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/kalogs-c/web-socrates/util"
)

type JWTCustomClaims struct {
	jwt.RegisteredClaims
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

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

func (s *service) Login(c context.Context, req *LoginRequest) (*LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.r.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(time.Hour * 24 * 7))),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		Token: tokenString,
		User: UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Lastname: user.Lastname,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	return response, nil
}
