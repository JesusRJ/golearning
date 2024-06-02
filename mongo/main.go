package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JesusRJ/golearning/mongo/codecs"
	"github.com/JesusRJ/golearning/mongo/dynamic"
	"github.com/JesusRJ/golearning/mongo/find"
	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/JesusRJ/golearning/mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Opts struct {
	Insert bool
}

const mongDSN = "mongodb://root:MongoPass321!@localhost:27017"
const Global = "Global"

var opts Opts = Opts{}

func init() {
	flag.BoolVar(&opts.Insert, "a", false, "add new documents in collections")
	flag.Parse()
}

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

	// dynamic.InsertNewUser(ctx, collUsers)

	now := time.Now().Format("04-05")

	newUser := model.User{
		Name: fmt.Sprintf("Dynamic Test %v", now),
		Company: &model.Company{
			Entity: model.Entity{ID: utils.GetCompanyID()},
		},
		Address: &model.Address{Street: fmt.Sprintf("street %v", now), Number: 123},
		Pets: []*model.Pet{
			{Name: "Dog"},
			{Name: "Cat"},
		},
	}

	if opts.Insert {
		dynamic.InsertNewUserByRef(ctx, collUsers, &newUser)
	}

	// find.Find(ctx, collUsers)
	fmt.Println("Users")
	find.FindWithAggregate[model.User](ctx, collUsers, filterUserByCompany(utils.GetCompanyID()))

	collPets := db.Collection(model.CollPets)
	fmt.Println("Pets")
	find.FindWithAggregate[model.Pet](ctx, collPets, filterPetsByUser(utils.GetUserID()))
}

func filterUserByCompany(companyID any) find.Filter {
	return func() bson.D {
		return bson.D{{Key: "$match", Value: bson.M{"company_id": companyID}}}
	}
}

func filterPetsByUser(userID any) find.Filter {
	return func() bson.D {
		return bson.D{{Key: "$match", Value: bson.M{"user_id": userID}}}
	}
}
