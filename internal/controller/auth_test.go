package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	controller := NewAuthController()
	userDto := dto.User{
		Username: "test",
		Password: "test",
	}

	t.Run("SignUp error when send bad userDto because email was required", func(t *testing.T) {

		userJSON, err := json.Marshal(userDto)
		assert.NoError(t, err)
		body := ioutil.NopCloser(bytes.NewBuffer(userJSON))

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		ginContext.Request = httptest.NewRequest(http.MethodGet, "/auth/signup", body)

		controller.SignUp(&ginContext)

		assert.Equal(t, http.StatusBadRequest, ginContext.Writer.Status())
	})
}
