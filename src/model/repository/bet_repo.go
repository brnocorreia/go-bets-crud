package repository

import (
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MONGODB_BET_COLLECTION = "MONGODB_BET_COLLECTION"
)

func NewBetRepository(
	database *mongo.Database,
) BetRepository {
	return &betRepository{
		database,
	}
}

type betRepository struct {
	databaseConnection *mongo.Database
}

type BetRepository interface {
	CreateBet(
		betDomain model.BetDomainInterface,
	) (model.BetDomainInterface, *rest_err.RestErr)
}
