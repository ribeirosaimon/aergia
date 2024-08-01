package controller

import (
	"github.com/gin-gonic/gin"
)

type HealthControllerInterface interface {
	GetHealth(c *gin.Context)
}

type AuthControllerInterface interface {
	Login(c *gin.Context)
}
