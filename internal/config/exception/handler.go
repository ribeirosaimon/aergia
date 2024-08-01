package exception

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AergiaError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e *AergiaError) Error() string {
	return e.Message
}

func BadRequest() *AergiaError {
	return &AergiaError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Time:    time.Now(),
	}
}

func (e *AergiaError) HandleError(c *gin.Context) {
	c.AbortWithStatusJSON(e.Code, e)
	c.Abort()
}
