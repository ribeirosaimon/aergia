package service

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestServiceAuth(t *testing.T) {
	service := newAuthServiceImpl()
	ctx := context.Background()

	for _, u := range []struct {
		testName string
		user     dto.User
		asError  bool
		errorMsg string
	}{
		{
			testName: "Error Without Password",
			asError:  true,
			user:     dto.User{Email: "user@example.com"},
			errorMsg: "password required",
		},
		{
			testName: "Email is not valid",
			asError:  true,
			user:     dto.User{Email: "userexample.com"},
			errorMsg: "password required",
		},
		{
			testName: "Error Without Email",
			asError:  true,
			user:     dto.User{Password: "AbcdEfh7!"},
			errorMsg: "email required",
		},
	} {
		t.Run(u.testName, func(t *testing.T) {
			err := service.SignUp(ctx, &u.user)
			if u.asError {
				assert.NotNil(t, err)
				assert.Equal(t, u.errorMsg, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
