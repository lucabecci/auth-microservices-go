package tests_test

import (
	"context"
	"testing"

	"github.com/lucabecci/auth-microservices-go/user"
)

func TestMakeValidateUserEndpoint(t *testing.T) {
	s := user.NewService()
	endpoint := user.MakeValidateUserEndpoint(s)
	//valid
	t.Run("valid user", func(t *testing.T) {
		req := user.ValidateUserRequest{
			Email:    "test@test.com",
			Password: "test",
		}
		_, err := endpoint(context.Background(), req)
		if err != nil {
			t.Errorf("Expected %v received %v", nil, err)
		}
	})
	//invalid
	t.Run("invalid user", func(t *testing.T) {
		req := user.ValidateUserRequest{
			Email:    "bad@bad.com",
			Password: "bad",
		}
		_, err := endpoint(context.Background(), req)
		if err == nil {
			t.Errorf("expected %v received %v", user.ErrInvalidUser, err)
		}
	})
}
