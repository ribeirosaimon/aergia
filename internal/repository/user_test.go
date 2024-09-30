package repository

import (
	"context"
	"testing"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/testutils/aergiatestcontainer"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	ctx := context.Background()
	pgsqlUrl, err := aergiatestcontainer.Pgsql(ctx)
	assert.Nil(t, err)

	err = sql.MockCreateTableDatabase(pgsqlUrl, map[string]bool{
		"user.sql": true,
	})
	assert.Nil(t, err)

	properties.NewMockPropertiesFile(map[string]string{
		"postgress.url":          pgsqlUrl,
		"postgress.database":     "postgres",
		string(constants.AERGIA): string(constants.DEV),
	})

	repository := NewUserRepository()

	t.Run("insert user in database", func(t *testing.T) {
		user, err := sql.NewUser(
			"user",
			"P@sw0rd!",
			"test@test.com",
			"firstName",
			"lastName",
		)
		assert.Nil(t, err)

		err = repository.InsertUser(ctx, user)
		assert.Nil(t, err)
	})
}
