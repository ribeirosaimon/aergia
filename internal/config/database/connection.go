package database

import (
	"context"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/mongo"
	"github.com/ribeirosaimon/aergia-utils/properties"
)

func NewConnection(ctx context.Context) mongo.AergiaMongoInterface {
	return mongo.NewConnMongo(ctx, mongo.WithUrl(properties.GetEnvironmentValue(constants.MongoProperties)))
}
