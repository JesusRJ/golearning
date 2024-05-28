package dynamic

import (
	"context"
	"reflect"

	"github.com/JesusRJ/golearning/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertNewUser(ctx context.Context, coll *mongo.Collection) *model.User {
	fields := []reflect.StructField{
		{Name: "Name", Type: reflect.TypeOf(""), Tag: reflect.StructTag(`bson:"name"`)},
		{Name: "Company", Type: reflect.TypeOf(primitive.NilObjectID), Tag: reflect.StructTag(`bson:"company_id"`)},
	}

	defType := reflect.StructOf(fields)

	value := reflect.New(defType)
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")

	value.Elem().FieldByName("Name").SetString("Dynamic Test")
	value.Elem().FieldByName("Company").Set(reflect.ValueOf(id))

	coll.InsertOne(ctx, value.Interface())

	return nil
}
