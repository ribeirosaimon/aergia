package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/aergia/internal/service"
)

var authGroup = "auth"
var authOnce sync.Once
var authController AuthControllerInterface

func NewAuthController() AuthControllerInterface {
	authOnce.Do(func() {
		authController = newAuthControllerImpl()
	})
	return authController
}

type authControllerImpl struct {
	authService service.AuthServiceInterface
}

func newAuthControllerImpl() AuthControllerInterface {
	return &authControllerImpl{
		authService: service.NewAuthService(),
	}
}

func (a *authControllerImpl) Login(c *gin.Context) {
	type LoginRequest struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var loginReq LoginRequest

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		return
	}

	a.authService.Login(c, loginReq.Login, loginReq.Password)
}

func init() {
	NewAergiaController(authGroup, "", http.MethodPost, NewAuthController().Login)
}
