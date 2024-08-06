package entity

import (
	"github.com/ribeirosaimon/aergia-utils/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
}

func (u *User) SetId(id primitive.ObjectID) {
	u.Id = id
}

// GetId don´t use a pointer
func (u *User) GetId() primitive.ObjectID {
	return u.Id
}

// GetId don´t use a pointer
func NewUser() mongo.Entity {
	var user User
	return &user
}
