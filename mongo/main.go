package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/JesusRJ/golearning/mongo/codecs"
	"github.com/JesusRJ/golearning/mongo/dynamic"
	"github.com/JesusRJ/golearning/mongo/find"
	"github.com/JesusRJ/golearning/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongDSN = "mongodb://root:MongoPass321!@localhost:27017"

func main() {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(mongDSN).SetRegistry(codecs.CustomRegistry())

	Client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("error connecting to database: %+v", err)
		os.Exit(1)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed pinging database: %+v", err)
		os.Exit(1)
	}

	db := Client.Database("petshop_test")
	collUsers := db.Collection(model.CollUser)

	newUser := dynamic.InsertNewUser(ctx, collUsers)
	fmt.Printf("%+v", newUser)

	// find.Find(ctx, collUsers)
	find.FindWithAggregate(ctx, collUsers)
}
