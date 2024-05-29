package find

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/JesusRJ/golearning/mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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
	match := bson.D{{Key: "$match", Value: bson.M{"company_id": getCompanyID()}}}
	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			// "company_id/company" Value is bson name in Company field's User model reference.
			// Could be anything, but by convention and to save reference between model and document db,
			// prefer add suffix "_id" to end
			// {Key: "from", Value: "company"}, // first test
			{Key: "from", Value: "company_id"},
			{Key: "localField", Value: "company_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "companies"},
		}},
	}
	lookupPet := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "pet"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "user_id"},
			{Key: "as", Value: "pets"},
		}},
	}

	setFields := bson.D{
		// Here "company_id" should be like $lookup, keeping reference to User model Company field bson name
		// {Key: "$set", Value: bson.M{"company": bson.M{"$arrayElemAt": bson.A{"$companies", 0}}}}, // first test
		{Key: "$set", Value: bson.M{"company_id": bson.M{"$arrayElemAt": bson.A{"$companies", 0}}}},
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

func getCompanyID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")
	return id
}
