package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/JesusRJ/golearning/mongo/codecs"
	"github.com/JesusRJ/golearning/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongDSN = "mongodb://root:MongoPass321!@localhost:27017"

// var Database *mongo.Database

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

	db := Client.Database("petshop")
	collUsers := db.Collection(model.CollUser)

	// Find(ctx, collUsers)
	FindWithAggregate(ctx, collUsers)
}

func Find(ctx context.Context, coll *mongo.Collection) {
	c, err := coll.Find(ctx, bson.M{"company_id": getCompanyID()})
	if err != nil {
		log.Printf("failed to get users: %+v", err)
		os.Exit(1)
	}

	var users []model.User
	if err := c.All(ctx, &users); err != nil {
		log.Printf("failed to decode users: %+v", err)
		os.Exit(1)
	}

	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

func FindWithAggregate(ctx context.Context, coll *mongo.Collection) {
	match := bson.D{{"$match", bson.M{"company_id": getCompanyID()}}}
	lookup := bson.D{
		{"$lookup", bson.D{
			{"from", "company"},
			{"localField", "company_id"},
			{"foreignField", "_id"},
			{"as", "companies"},
		}},
	}
	// set := bson.D{{"$unwind", bson.M{
	// 	"path":                       "$companies",
	// 	"preserveNullAndEmptyArrays": true,
	// }}}

	addFields := bson.D{
		{"$addFields", bson.M{"company": bson.M{"$arrayElemAt": bson.A{"$companies", 0}}}},
	}

	c, err := coll.Aggregate(ctx, mongo.Pipeline{match, lookup, addFields})
	if err != nil {
		log.Printf("failed to get users (aggregate): %+v", err)
		os.Exit(1)
	}

	var users []model.User
	if err := c.All(ctx, &users); err != nil {
		log.Printf("failed to decode users: %+v", err)
		os.Exit(1)
	}

	for _, u := range users {
		// fmt.Printf("%+v | %s \n", u, u.Company[0].Name)
		fmt.Printf("%+v", u)
	}
}

func getCompanyID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")
	return id
}
