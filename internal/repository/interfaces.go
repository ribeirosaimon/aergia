package repository

import (
	"context"

	"github.com/ribeirosaimon/aergia-utils/entities/sql"
)

type UserRepositoryInterface interface {
	GetUser(ctx context.Context, id string) (*sql.User, error)
}
