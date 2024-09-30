package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/response"
	"github.com/ribeirosaimon/aergia/internal/controller/mock"
	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/ribeirosaimon/aergia/internal/service"
)

type AuthControllerInterface interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

var authGroup = "auth"
var authOnce sync.Once
var authController AuthControllerInterface

func NewAuthController() AuthControllerInterface {
	authOnce.Do(func() {
		switch properties.GetEnvironmentMode() {
		case constants.PROD, constants.DEV, constants.INTEGRATION:
			authController = newAuthControllerImpl()
		default:
			authController = new(mock.AuthControllerMock)
		}
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

func (a *authControllerImpl) SignUp(c *gin.Context) {
	var userDto dto.User
	if err := c.ShouldBindJSON(&userDto); err != nil {
		response.AergiaResponseStatusBadRequest(c, err)
		return
	}
	if err := a.authService.SignUp(c, &userDto); err != nil {
		response.AergiaResponseStatusBadRequest(c, err)
	}
	response.AergiaResponseOk(c, nil)
}

func (a *authControllerImpl) Login(c *gin.Context) {

	var loginReq dto.Login

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.AergiaResponseStatusBadRequest(c, err)
	}

	if err := a.authService.Login(c, loginReq.Email, loginReq.Password); err != nil {
		response.AergiaResponseStatusBadRequest(c, err)
	}
}

func AuthControllers() {
	NewAergiaController(authGroup, "", http.MethodPost, NewAuthController().Login)
	NewAergiaController(authGroup, "/signup", http.MethodPost, NewAuthController().SignUp)
}
