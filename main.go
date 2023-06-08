package main

import (
	"context"
	"log"

	"github.com/brnocorreia/go-movies-crud/src/configuration/database/mongodb"
	"github.com/brnocorreia/go-movies-crud/src/configuration/logger"
	"github.com/brnocorreia/go-movies-crud/src/controller"
	"github.com/brnocorreia/go-movies-crud/src/controller/routes"
	"github.com/brnocorreia/go-movies-crud/src/model/repository"
	"github.com/brnocorreia/go-movies-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
	}

	// Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
