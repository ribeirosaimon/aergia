package repository

import (
	"context"
	"testing"
	"time"

	"github.com/ribeirosaimon/aergia-utils/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	ctx := context.Background()
	// pgsqlUrl, err := aergiatestcontainer.Pgsql(ctx)
	// assert.Nil(t, err)

	pgsqlUrl := "jdbc:postgresql://localhost:5432/postgres?user=postgres&password=frajolinha202"

	properties.NewMockPropertiesFile(map[string]string{
		"postgress.url":      pgsqlUrl,
		"postgress.database": "postgres",
	})
	// connPgsql := pgsql.NewConnPgsql(pgsql.WithUrl(pgsqlUrl), pgsql.WithDatabase("motion"))
	repository := NewUserRepository()

	t.Run("insert user in database", func(t *testing.T) {
		user := sql.User{
			Password:  "password",
			Username:  "username",
			Email:     "email",
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
