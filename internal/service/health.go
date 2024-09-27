package service

import (
	"os"
	"sync"
	"time"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia/internal/dto"
)

type HealthService interface {
	GetHealth() (*dto.Health, error)
}

var helathOnce sync.Once
var healthService HealthService

// NewHealthService is once open function
func NewHealthService() HealthService {
	helathOnce.Do(func() {
		healthService = newHealthServiceImpl()
	})
	return healthService
}

type healthServiceImpl struct {
}

func newHealthServiceImpl() HealthService {
	return &healthServiceImpl{}
}

func (h *healthServiceImpl) GetHealth() (*dto.Health, error) {
	return &dto.Health{
		Status:      "up",
		Date:        time.Now(),
		Environment: os.Getenv(string(constants.AERGIA)),
	}, nil
}
