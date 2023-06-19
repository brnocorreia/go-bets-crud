package repository

import (
	"context"
	"fmt"
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository/entity"
	"github.com/brnocorreia/go-movies-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (br *betRepository) FindBetByID(
	id string,
) (model.BetDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findBetByID repository",
		zap.String("journey", "findBetByID"))

	collection_name := os.Getenv(MONGODB_BET_COLLECTION)
	collection := br.databaseConnection.Collection(collection_name)

	betEntity := &entity.BetEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(betEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Bet not found with this id: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findBetByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find bet by id."
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findBetByID"))
		return nil, rest_err.NewInternalServerError(errorMessage, []rest_err.Causes{})
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", betEntity.ID.Hex()),
	)
	return converter.ConvertBetEntityToDomain(*betEntity), nil
}
