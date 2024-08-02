package service

import (
	"context"
	"sync"

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

func (a *authServiceImpl) Login(ctx context.Context, login, pass string) error {

}
