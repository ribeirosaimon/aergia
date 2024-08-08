package repository

import (
	"context"
	"sync"

	"github.com/ribeirosaimon/aergia-utils/entities/sql"
	"github.com/ribeirosaimon/aergia/internal/config/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userOnce sync.Once
var userRepository UserRepositoryInterface
var userTable = "user"

// NewUserRepository is once open function
func NewUserRepository(ctx context.Context) UserRepositoryInterface {
	userOnce.Do(func() {
		userRepository = newUserRepositoryImpl(ctx)
	})
	return userRepository
}

func newUserRepositoryImpl(ctx context.Context) UserRepositoryInterface {
	return &UserRepositoryImpl{
		conn: database.NewConnection(ctx).GetConnection().Collection(userTable),
	}
}

type UserRepositoryImpl struct {
	conn *mongo.Collection
}

func (u *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*sql.User, error) {
	var user sql.User
	if err := u.conn.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
