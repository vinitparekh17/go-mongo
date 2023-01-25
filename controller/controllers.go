package controller

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var MongoClient *mongo.Client

func Init() {
	err := godotenv.Load()
	fetal(err)
	clientOpt := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	MongoClient, err = mongo.Connect(context.TODO(), clientOpt)
	fetal(err)
	DisconnectMongo(MongoClient)
	Collection = MongoClient.Database("testing").Collection("go-movie")
}

func DisconnectMongo(Mongo *mongo.Client) {
	err := Mongo.Disconnect(context.Background())
	fetal(err)
}

func fetal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
