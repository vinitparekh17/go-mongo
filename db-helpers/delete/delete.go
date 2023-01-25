package delete

import (
	"context"
	"fmt"
	"go-mongo/controller"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func deleteOne(movideId string) {
	id, _ := primitive.ObjectIDFromHex(movideId)
	filter := bson.M{"_id": id}
	res, _ := controller.Collection.DeleteOne(context.Background(), filter)

	fmt.Println(res)
}

func deleteAll() {
	filter := bson.D{{}}
	res, err := controller.Collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.DeletedCount)
}
