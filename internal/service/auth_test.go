package service

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestServiceSignUp(t *testing.T) {
	service := newAuthServiceImpl()
	ctx := context.Background()

	for _, u := range []struct {
		testName string
		user     dto.User
		asError  bool
		errorMsg string
	}{
		{
			testName: "Success create user",
			asError:  false,
			user:     dto.User{Email: "user@example.com", Password: "P@sw0rd!"},
		},
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
			errorMsg: "invalid email",
		},
		{
			testName: "Error Without Email",
			asError:  true,
			user:     dto.User{Password: "AbcdEfh7!"},
			errorMsg: "invalid email",
		},
		{
			testName: "Error Without Password incorrectly",
			asError:  true,
			user:     dto.User{Email: "test@teste.com", Password: "123!"},
			errorMsg: "password too short",
		},
		{
			testName: "Error Password incorrectly without uppercase",
			asError:  true,
			user:     dto.User{Email: "test@teste.com", Password: "asdasdasdasd1@!"},
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
		},
		{
			testName: "Error Password incorrectly without lowercase",
			asError:  true,
			user:     dto.User{Email: "test@teste.com", Password: "ASDASDASDASD@!"},
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
		},
		{
			testName: "Error Password incorrectly without special character",
			asError:  true,
			user:     dto.User{Email: "test@teste.com", Password: "asdasdasdASd123"},
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
		},
		{
			testName: "Error Password incorrectly without number",
			asError:  true,
			user:     dto.User{Email: "test@teste.com", Password: "AbcdEfh!as"},
			errorMsg: "password must contain uppercase and lower case, digit and special characters",
		},
		{
			testName: "Error Without valid email",
			asError:  true,
			user:     dto.User{Email: "testteste.com", Password: "Abcd12Efh!as"},
			errorMsg: "invalid email",
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

func TestServiceLogin(t *testing.T) {
	service := newAuthServiceImpl()
	ctx := context.Background()

	for _, u := range []struct {
		testName string
		user     dto.Login
		asError  bool
		errorMsg string
	}{
		{
			testName: "Success login",
			user:     dto.Login{Email: "test@teste.com", Password: "Abcd12Efh!as"},
		},
	} {
		t.Run(u.testName, func(t *testing.T) {
			err := service.Login(ctx, u.user.Email, u.user.Password)
			if u.asError {
				assert.NotNil(t, err)
				assert.Equal(t, u.errorMsg, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
