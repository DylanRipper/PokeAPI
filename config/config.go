package config

import (
	"context"
	"pokemon/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(c context.Context) (r *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err := client.Connect(ctx); err != nil {
		panic(err)
	}
	r = client.Database("pokemon")
	return r
}

var DB *mongo.Database

func Query(db *mongo.Database) *model.QueryMongo {
	return &model.QueryMongo{
		Db:         db,
		Collection: "user",
	}
}
