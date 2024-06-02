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

type Filter func() bson.D

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

func FindWithAggregate[T any](ctx context.Context, coll *mongo.Collection, filter Filter) {
	t := reflect.TypeOf(new(T))
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// The first element should be the match filter
	// ex.: bson.D{{Key: "$match", Value: bson.M{"user_id": userID}}}
	pipeline := []bson.D{filter()}

	for x := range t.NumField() {
		field := t.Field(x)
		st, err := parser.DefaultStructTagParser(field)
		if err != nil {
			panic(err)
		}

		switch st.Relation {
		case parser.BelongsTo:
			pipeline = append(pipeline, createLookup(st), createSetField(st))
		case parser.HasMany:
			pipeline = append(pipeline, createLookup(st))
		}
	}

	c, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Printf("failed to get entity (aggregate): %+v", err)
		os.Exit(1)
	}

	var users []T
	if err := c.All(ctx, &users); err != nil {
		log.Printf("failed to decode entity: %+v", err)
		os.Exit(1)
	}

	utils.Print(users)
}

func createLookup(st parser.StructTag) bson.D {
	return bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: st.From},                 // collection to join
			{Key: "localField", Value: st.LocalField},     // field from the input documents
			{Key: "foreignField", Value: st.ForeignField}, // field from the documents of the "from" collection
			{Key: "as", Value: st.As},                     // output array field
		}},
	}
}

func createSetField(st parser.StructTag) bson.D {
	return bson.D{
		{Key: "$set", Value: bson.M{st.As: bson.M{"$arrayElemAt": bson.A{fmt.Sprintf("$%s", st.As), 0}}}},
	}
}
