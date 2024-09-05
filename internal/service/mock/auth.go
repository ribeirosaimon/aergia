package mock

import (
	"context"

	"github.com/ribeirosaimon/aergia/internal/dto"
)

type AuthServiceMock struct {
}

func (a *AuthServiceMock) Login(ctx context.Context, login, pass string) error {
	// TODO implement me
	panic("implement me")
}

func (a *AuthServiceMock) SignUp(ctx context.Context, d *dto.User) error {
	// TODO implement me
	panic("implement me")
}
