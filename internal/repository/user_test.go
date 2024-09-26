package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

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

	// melhor ser string, pois esse valor tem que ser imutavel
	properties.NewMockPropertiesFile(map[string][]byte{
		"postgress.url":          []byte(pgsqlUrl),
		"postgress.database":     []byte("postgres"),
		string(constants.AERGIA): []byte(constants.DEV),
	})

	repository := NewUserRepository()

	t.Run("insert user in database", func(t *testing.T) {

		user := sql.User{
			Password:  "password",
			Username:  fmt.Sprintf("username_%d", 1),
			Email:     fmt.Sprintf("email_%d", 1),
			FirstName: "first_name",
			LastName:  "last_name",
			Role:      "role",
			Audit: sql.Audit{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		createUser, err := repository.CreateUser(ctx, &user)
		assert.Nil(t, err)
		assert.NotNil(t, createUser)
	})
}
