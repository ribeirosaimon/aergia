package database

import (
	"context"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/storage/mongo"
	"github.com/ribeirosaimon/aergia-utils/storage/pgsql"
)

func NewMongoConnection(ctx context.Context) mongo.AergiaMongoInterface {
	return mongo.NewConnMongo(ctx,
		mongo.WithUrl(string(properties.GetEnvironmentValue(constants.MongoUrl))),
		mongo.WithDatabase(string(properties.GetEnvironmentValue(constants.MongoDatabase))),
	)
}
func NewPgsqlConnection() pgsql.AergiaPgsqlInterface {
	return pgsql.NewConnPgsql(
		pgsql.WithUrl(string(properties.GetEnvironmentValue(constants.PostgressUrl))),
	)
}
