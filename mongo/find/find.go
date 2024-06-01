package find

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/JesusRJ/golearning/mongo/parser"
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

func FindWithAggregate[T any](ctx context.Context, coll *mongo.Collection) {
	tType := reflect.TypeOf(new(T))
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}

	// Filter
	match := bson.D{{Key: "$match", Value: bson.M{"company_id": utils.GetCompanyID()}}}

	pipeline := []bson.D{match}
	for x := range tType.NumField() {
		field := tType.Field(x)
		st, err := parser.DefaultStructTagParser(field)
		if err != nil {
			panic(err)
		}

		if st.BelongsTo() || st.HasMany() {
			lookup := bson.D{
				{Key: "$lookup", Value: bson.D{
					{Key: "from", Value: st.From},                 // collection to join
					{Key: "localField", Value: st.LocalField},     // field from the input documents
					{Key: "foreignField", Value: st.ForeignField}, // field from the documents of the "from" collection
					{Key: "as", Value: st.As},                     // output array field
				}},
			}
			pipeline = append(pipeline, lookup)
		}
	}

	setFields := bson.D{
		// Here "company" should be like $lookup, keeping reference to User model Company field bson name
		{Key: "$set", Value: bson.M{"company": bson.M{"$arrayElemAt": bson.A{"$companies", 0}}}},
	}

	pipeline = append(pipeline, setFields)

	c, err := coll.Aggregate(ctx, pipeline)
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
