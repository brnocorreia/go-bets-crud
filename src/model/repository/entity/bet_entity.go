package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BetEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Nickname    string             `bson:"nickname,omitempty"`
	Date        string             `bson:"date,omitempty"`
	Bookmaker   string             `bson:"bookmaker,omitempty"`
	Sport       string             `bson:"sport,omitempty"`
	Description string             `bson:"description,omitempty"`
	Odd         float32            `bson:"odd,omitempty"`
	Investment  float32            `bson:"investment,omitempty"`
	Returned    float32            `bson:"return,omitempty"`
	Profit      float32            `bson:"profit,omitempty"`
	Tipster     string             `bson:"tipster,omitempty"`
	Winner      bool               `bson:"winner,omitempty"`
}
