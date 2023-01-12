package main

import (
	"context"
	"log"

	"github.com/Avinashanakal/controllers"
	"github.com/Avinashanakal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	ctx            context.Context
	mongoClient    *mongo.Client
	userCollection *mongo.Collection
	userService    services.UserService
	UserController controllers.UserController
	err            error
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Panic(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	userCollection = mongoClient.Database("userdb").Collection("users")
	userService = services.NewUser(userCollection, ctx)
	UserController = controllers.New(userService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basePath := server.Group("/v1")
	UserController.RegisterRoutes(basePath)

	log.Fatal(server.Run(":8000"))
}
