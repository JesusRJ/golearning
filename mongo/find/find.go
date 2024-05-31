package find

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/JesusRJ/golearning/mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(ctx context.Context, coll *mongo.Collection) {
	c, err := coll.Find(ctx, bson.M{"company_id": utils.GetCompanyID()})
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
	match := bson.D{{Key: "$match", Value: bson.M{"company_id": utils.GetCompanyID()}}}
	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "company"},          // collection to join
			{Key: "localField", Value: "company_id"}, // field from the input documents
			{Key: "foreignField", Value: "_id"},      // field from the documents of the "from" collection
			{Key: "as", Value: "companies"},          // output array field
		}},
	}
	lookupPet := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "pet"},             // collection to join
			{Key: "localField", Value: "_id"},       // field from the input documents
			{Key: "foreignField", Value: "user_id"}, // field from the documents of the "from" collection
			{Key: "as", Value: "pets"},              // output array field
		}},
	}

	setFields := bson.D{
		// Here "company" should be like $lookup, keeping reference to User model Company field bson name
		{Key: "$set", Value: bson.M{"company": bson.M{"$arrayElemAt": bson.A{"$companies", 0}}}},
	}

	c, err := coll.Aggregate(ctx, mongo.Pipeline{match, lookup, lookupPet, setFields})
	if err != nil {
		log.Printf("failed to get users (aggregate): %+v", err)
		os.Exit(1)
	}

	var users []model.User
	if err := c.All(ctx, &users); err != nil {
		log.Printf("failed to decode users: %+v", err)
		os.Exit(1)
	}

	utils.PrintUsers(users)
}
