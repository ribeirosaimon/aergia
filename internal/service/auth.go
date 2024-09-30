package service

import (
	"context"
	"sync"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/ribeirosaimon/aergia/internal/repository"
	"github.com/ribeirosaimon/aergia/internal/service/mock"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, login, pass string) error
	SignUp(ctx context.Context, d *dto.User) error
}

var authOnce sync.Once
var authService AuthServiceInterface

// NewAuthService is once open function
func NewAuthService() AuthServiceInterface {
	authOnce.Do(func() {
		switch properties.GetEnvironmentMode() {
		case constants.PROD, constants.DEV, constants.INTEGRATION:
			authService = newAuthServiceImpl()
		default:
			authService = new(mock.AuthServiceMock)
		}

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
	newUser, err := sql.NewUser(user.Username, user.Password, user.Email, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	if err = a.userRepository.InsertUser(ctx, newUser); err != nil {
		return err
	}

	return nil
}

func (a *authServiceImpl) Login(ctx context.Context, login, pass string) error {
	if err := a.userRepository.GetUser(ctx, newUser); err != nil {
		return err
	}

}
