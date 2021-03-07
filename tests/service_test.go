package tests_test

import (
	"context"
	"testing"

	"github.com/lucabecci/auth-microservices-go/user"
	"github.com/stretchr/testify/assert"
)

func TestValidateTest(t *testing.T) {
	service := user.NewService()
	t.Run("invalid user", func(t *testing.T) {
		_, err := service.ValidateUser(context.Background(), "bad@test.com", "bad")
		assert.NotNil(t, err)
		assert.Equal(t, "User, invalid", err.Error())
	})
	t.Run("valid user", func(t *testing.T) {
		token, err := service.ValidateUser(context.Background(), "test@test.com", "test")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
}
