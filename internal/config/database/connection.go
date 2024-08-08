package database

import (
	"context"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia-utils/storage/mongo"
)

func NewConnection(ctx context.Context) mongo.AergiaMongoInterface {
	return mongo.NewConnMongo(ctx, mongo.WithUrl(properties.GetEnvironmentValue(constants.MongoProperties)))
}
