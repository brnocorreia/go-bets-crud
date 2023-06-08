package repository

import (
	"context"
	"os"

	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/configuration/rest_err"
	"github.com/brnocorreia/go-movies-crud/src/model"
	"github.com/brnocorreia/go-movies-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error(), []rest_err.Causes{})
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*value), nil
}
