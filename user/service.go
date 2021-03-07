package user

import (
	"context"
	"errors"
)

type Service interface {
	ValidateUser(ctx context.Context, email string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

var (
	ErrInvalidUser  = errors.New("User invalid")
	ErrInvalidToken = errors.New("Token invalid")
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) ValidateUser(ctx context.Context, email string, password string) (string, error) {
	if email != "test@test.com" && password != "test" {
		return "nil", ErrInvalidUser
	}

	// token, err := security.NewToken(email)
	// if err != nil {
	//     return "", err
	// }

	return "", nil
}

func (s *service) ValidateToken(ctx context.Context, token string) (string, error) {
	// t, err := security.ParseToken(token)
	// if err != nil {
	// 	return "", ErrInvalidToken
	// }

	// tData, err := security.GetClaims(t)
	// if err != nil {
	// 	return "", ErrInvalidToken
	// }
	// return tData["email"].(string), nil
	return "", nil
}
