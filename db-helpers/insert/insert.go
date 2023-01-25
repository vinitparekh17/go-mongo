package insert

// importing / exporting packages
import (
	"context"
	"go-mongo/controller"
	"go-mongo/model"
)

func insertOne(movie model.Netflix) {
	controller.Init()
	res, _ := controller.Collection.InsertOne(context.Background(), movie)
	println(res)
	controller.DisconnectMongo()
}
