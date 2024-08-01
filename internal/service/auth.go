package service

import (
	"context"
	"sync"
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
}

func newAuthServiceImpl() AuthServiceInterface {
	return &authServiceImpl{}
}

func (a authServiceImpl) Login(ctx context.Context, login, pass string) error {
	// TODO implement me
	panic("implement me")
}
