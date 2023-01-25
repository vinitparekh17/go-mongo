package update

import (
	"context"
	"go-mongo/controller"
	"go-mongo/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateOne(movieId string, movie model.Netflix) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"title": "godzilla"}}

	res, err := controller.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	println(res.ModifiedCount)
}
