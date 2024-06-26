package dynamic

import (
	"context"
	"fmt"
	"reflect"

	"github.com/JesusRJ/golearning/mongo/model"
	"github.com/JesusRJ/golearning/mongo/parser"
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

	fields := []reflect.StructField{}
	values := []any{}
	for x := range v.NumField() {
		valueField := v.Field(x)
		typeField := v.Type().Field(x)

		structTag, err := parser.DefaultStructTagParser(typeField)
		if err != nil {
			panic(err)
		}

		value := valueField.Interface()

		switch structTag.Relation {
		case parser.BelongsTo:
			// Convert field to primitive.ObjectID
			typeField = reflect.StructField{
				Name: typeField.Name,
				Type: reflect.TypeOf(primitive.NilObjectID),
				Tag:  reflect.StructTag(fmt.Sprintf(`bson:"%s"`, structTag.LocalField)),
			}
			value = valueField.Interface().(model.AbstractEntity).GetID()
		case parser.HasMany:
			// Ignore hasMany field on insert/update operation
			continue
		}

		fields = append(fields, typeField)
		values = append(values, value)
	}

	defType := reflect.StructOf(fields)
	value := reflect.New(defType)

	for v := range values {
		value.Elem().Field(v).Set(reflect.ValueOf(values[v]))
	}

	coll.InsertOne(ctx, value.Interface())
}
