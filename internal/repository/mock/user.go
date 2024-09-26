package mock

import (
	"context"

	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
)

type UserRepositoryMock struct {
}

func (u UserRepositoryMock) CreateUser(ctx context.Context, user *sql.User) (*sql.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserRepositoryMock) GetUser(ctx context.Context, id string) (*sql.User, error) {
	// TODO implement me
	panic("implement me")
}
