package get

import (
	"context"
	"fmt"
	"go-mongo/controller"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func getAll() {
	filter := bson.D{{}}
	cur, err := controller.Collection.Find(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	fmt.Println(results)
}
