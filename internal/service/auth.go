package service

import (
	"context"
	"sync"

	"github.com/ribeirosaimon/aergia-utils/entities/role"
	"github.com/ribeirosaimon/aergia-utils/entities/sql"
	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/ribeirosaimon/aergia/internal/repository"
)

var authOnce sync.Once
var authService AuthServiceInterface

// NewAuthService is once open function
func NewAuthService() AuthServiceInterface {
	authOnce.Do(func() {
		authService = newAuthServiceImpl()
	})
	return authService
}

type authServiceImpl struct {
	userRepository repository.UserRepositoryInterface
}

func newAuthServiceImpl() AuthServiceInterface {
	return &authServiceImpl{
		userRepository: repository.NewUserRepository(),
	}
}

func (a *authServiceImpl) SignUp(ctx context.Context, user *dto.User) error {
	var dbUser sql.User

	dbUser.Email = user.Email
	dbUser.Username = user.Username
	dbUser.Password = user.Password

	dbUser.Role = role.USER
	dbUser.LastName = user.LastName
	dbUser.FirstName = user.FirstName
	dbUser.LoginAtempt = 0

	_, err := a.userRepository.CreateUser(ctx, &dbUser)
	if err != nil {
		return err
	}
	return nil
}

func (a *authServiceImpl) Login(ctx context.Context, login, pass string) error {
	return nil
}
