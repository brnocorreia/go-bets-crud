package repository

import (
	"context"
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (br *betRepository) CreateBet(
	betDomain model.BetDomainInterface,
) (model.BetDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_BET_COLLECTION)

	collection := br.databaseConnection.Collection(collection_name)

	value := converter.ConvertBetDomainToEntity(betDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create bet",
			err,
			zap.String("journey", "createBet"))
		return nil, rest_err.NewInternalServerError(err.Error(), []rest_err.Causes{})
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("betUser repository executed successfully",
		zap.String("betId", value.ID.Hex()),
		zap.String("journey", "betUser"))

	return converter.ConvertBetEntityToDomain(*value), nil
}
