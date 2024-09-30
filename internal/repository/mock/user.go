package mock

import (
	"context"

	"github.com/pkg/errors"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) InsertUser(ctx context.Context, user *sql.User) error {
	called := u.Mock.MethodCalled("InsertUser", ctx, user)
	return called.Error(0)
}

func (u *UserRepositoryMock) FindUserByEmail(ctx context.Context, email string) (*sql.User, error) {
	called := u.Mock.MethodCalled("FindUserByEmail", ctx, email)
	if len(called) >= 1 {
		return called[0].(*sql.User), nil
	}
	return nil, errors.New("error called FindUserByEmail")
}
