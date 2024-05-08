package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/JesusRJ/golearning/mongo/codecs"
	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	match := bson.D{{Key: "$match", Value: bson.M{"company_id": getCompanyID()}}}
	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "company"},
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

	// for _, u := range users {
	// 	// fmt.Printf("%+v | %s \n", u, u.Company[0].Name)
	// 	fmt.Printf("%+v", u.Company)
	// }

	printUsers(users)
}

func getCompanyID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")
	return id
}

func printUsers(users []model.User) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"User", "Address", "Phones", "Company", "Pets"})

	for _, u := range users {
		phones := getCommaSeparated(u.Phone)
		pets := getCommaSeparated(u.Pets)

		t.AppendRows([]table.Row{
			{u.Name, u.Address, phones, u.Company, pets},
		})
	}

	t.AppendFooter(table.Row{"", "", "", "Count", len(users)})
	t.Render()
}

func getCommaSeparated[T interface{ String() string }](values []T) (result string) {
	for _, v := range values {
		result = fmt.Sprintf("%s, %s", result, v)
	}
	if result != "" {
		return result[1:]
	}
	return result
}
