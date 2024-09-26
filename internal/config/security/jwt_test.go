package security

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/role"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/util"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {

	t.Run("jwt process", func(t *testing.T) {
		properties.NewMockPropertiesFile(map[string][]byte{
			"secret_key": []byte("secret_key"),
		})

		loggedUser := &util.LoggerUser{
			Name: "testName",
			Role: role.USER,
		}
		token, err := CreateToken(loggedUser)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		ctx := context.Background()
		ctx, err = VerifyToken(ctx, token)

		ctxLoggedUser := ctx.Value(constants.LoggedUser).(util.LoggerUser)

		assert.NoError(t, err)
		assert.Equal(t, loggedUser, ctxLoggedUser)
	})
}
