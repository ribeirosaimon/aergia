package service

import (
	"context"

	"github.com/ribeirosaimon/aergia/internal/dto"
)

type HealthService interface {
	GetHealth() (*dto.Health, error)
}

type AuthServiceInterface interface {
	Login(ctx context.Context, login, pass string) error
	SignUp(ctx context.Context, d *dto.User) error
}
