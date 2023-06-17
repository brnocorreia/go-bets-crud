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

	// Init user dependencies
	repo_users := repository.NewUserRepository(database)
	service_users := service.NewUserDomainService(repo_users)
	userController := controller.NewUserControllerInterface(service_users)

	repo_bets := repository.NewBetRepository(database)
	service_bets := service.NewBetDomainService(repo_bets)
	betController := controller.NewBetControllerInterface(service_bets)

	router := gin.Default()

	routes.InitUserRoutes(&router.RouterGroup, userController)
	routes.InitBetRoutes(&router.RouterGroup, betController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
