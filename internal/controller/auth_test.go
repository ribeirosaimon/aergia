package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"github.com/ribeirosaimon/aergia/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	controller := NewAuthController()

	// all error in input dto
	for _, v := range []struct {
		testName string
		want     string
		userIput dto.User
	}{
		{testName: "error because only have username", want: "Email", userIput: dto.User{Username: "test"}},
		{testName: "error because only have email", want: "Username", userIput: dto.User{Email: "test"}},
		{testName: "error because only have email and username", want: "Username", userIput: dto.User{Password: "test"}},
	} {
		t.Run(v.testName, func(t *testing.T) {
			userJSON, err := json.Marshal(v.userIput)
			assert.NoError(t, err)
			body := ioutil.NopCloser(bytes.NewBuffer(userJSON))

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)
			c.Request = httptest.NewRequest(http.MethodGet, "/auth/signup", body)

			controller.SignUp(c)

			logs.LOG.Message(rr.Body.String())
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.True(t, strings.Contains(rr.Body.String(), v.want))
		})
	}

	t.Run("success signup", func(t *testing.T) {

		// pgsql, err := aergiatestcontainer.Pgsql(context.Background())
		// assert.NoError(t, err)

		userDto := dto.User{Username: "test", Email: "test@test.com", Password: "test"}

		userJSON, err := json.Marshal(userDto)
		assert.NoError(t, err)
		body := ioutil.NopCloser(bytes.NewBuffer(userJSON))

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Request = httptest.NewRequest(http.MethodGet, "/auth/signup", body)

		controller.SignUp(c)

		logs.LOG.Message(rr.Body.String())
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
