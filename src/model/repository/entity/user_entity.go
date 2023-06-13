package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Password    string             `bson:"password,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Age         int8               `bson:"age,omitempty"`
	Nationality string             `bson:"nationality,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
}
