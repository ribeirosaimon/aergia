package mock

import (
	"github.com/gin-gonic/gin"
)

type AuthControllerMock struct {
}

func (a *AuthControllerMock) Login(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (a *AuthControllerMock) SignUp(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}
