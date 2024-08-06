package repository

import (
	"context"

	"github.com/ribeirosaimon/aergia/internal/entity"
)

type UserRepositoryInterface interface {
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
