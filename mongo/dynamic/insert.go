package dynamic

import (
	"context"
	"reflect"

	"github.com/JesusRJ/golearning/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertNewUser(ctx context.Context, coll *mongo.Collection) {
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
}

func InsertNewUserByRef(ctx context.Context, coll *mongo.Collection, user *model.User) {
	v := reflect.ValueOf(user)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fields := make([]reflect.StructField, v.NumField())
	values := make([]any, v.NumField())
	for x := range v.NumField() {
		valueField := v.Field(x)
		typeField := v.Type().Field(x)

		var value any = valueField.Interface()
		if typeField.Tag.Get("ref") == "belongsTo" {
			// Convert company field to primitive.ObjectID
			typeField = reflect.StructField{
				Name: typeField.Name,
				Type: reflect.TypeOf(primitive.NilObjectID),
				Tag:  reflect.StructTag(`bson:"company_id"`),
			}

			value = valueField.Interface().(*model.Company).ID
		}
		fields[x] = typeField
		values[x] = value
	}

	defType := reflect.StructOf(fields)
	value := reflect.New(defType)

	for v := range values {
		value.Elem().Field(v).Set(reflect.ValueOf(values[v]))
	}

	coll.InsertOne(ctx, value.Interface())
}
