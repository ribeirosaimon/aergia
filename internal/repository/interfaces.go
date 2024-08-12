package repository

import (
	"context"

	"github.com/ribeirosaimon/aergia-utils/entities/sql"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *sql.User) (*sql.User, error)
	GetUser(ctx context.Context, id string) (*sql.User, error)
}
