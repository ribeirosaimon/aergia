package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/aergia/internal/service"
)

type HealthControllerInterface interface {
	GetHealth(c *gin.Context)
}

var healthGroup = "health"
var healthOnce sync.Once
var healthController HealthControllerInterface

func NewHealthController() HealthControllerInterface {
	healthOnce.Do(func() {
		healthController = newHealthControllerImpl()
	})
	return healthController
}

type healthControllerImpl struct {
	healthService service.HealthService
}

func newHealthControllerImpl() HealthControllerInterface {
	return &healthControllerImpl{
		healthService: service.NewHealthService(),
	}
}

func (h *healthControllerImpl) GetHealth(c *gin.Context) {
	health, err := h.healthService.GetHealth()
	if err != nil {
		panic(err)
	}
	c.JSON(200, health)
}

func init() {
	NewAergiaController(healthGroup, "", http.MethodGet, NewHealthController().GetHealth)
}
