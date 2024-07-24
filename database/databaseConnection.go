package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbinstance() *mongo.Client {
	mongodb_url := "mongodb+srv://udaysiddu492:password@cluster0.6ndv4e2.mongodb.net/"
	dbName:= "BookStoreManagement"
	client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI(mongodb_url+dbName))
	if err != nil {
		panic(err)
	}

	return client

}

var Client*mongo.Client= dbinstance()
var BookCollection *mongo.Collection= Client.Database("BookStoreManagement").Collection("BookCollection")
var OrderCollection *mongo.Collection= Client.Database("OrderStoreManagement").Collection("Collection")
var UserCollection *mongo.Collection= Client.Database("UserStoreManagement").Collection("UserCollection")