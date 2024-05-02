package main

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID string `bson:"x,ignore"`
}

func UserEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	fmt.Println("passou pelo Encoder")

	switch val.Kind() {
	case reflect.String:
		fmt.Println("É uma string")
	case reflect.Struct:
		fmt.Println("É uma struct")
	}
	return nil
}

func StringEncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	switch val.Kind() {
	case reflect.String:
		fmt.Printf("É uma string: %v", val.Interface().(string))
		vw.WriteString(fmt.Sprintf("É uma string: %v", val.Interface()))
	case reflect.Struct:
		fmt.Println("É uma struct")
	}
	return nil
}

func UserDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	fmt.Println("passou pelo Decoder")
	return nil
}

func main() {
	// Create a registry that decodes nothing.
	registry := bson.NewRegistry()

	registry.RegisterTypeEncoder(reflect.TypeOf(User{}), bsoncodec.ValueEncoderFunc(UserEncodeValue))
	// registry.RegisterTypeEncoder(reflect.TypeOf(string("")), bsoncodec.ValueEncoderFunc(StringEncodeValue))
	// registry.RegisterTypeDecoder(reflect.TypeOf(User{}), bsoncodec.ValueDecoderFunc(UserDecodeValue))

	client, err := mongo.Connect(context.Background(),
		options.Client().
			ApplyURI("mongodb://root:MongoPass321!@localhost:27017").
			SetRegistry(registry))

	if err != nil {
		panic(err)
	}

	coll := client.Database("test").Collection("coll")
	_, err = coll.InsertOne(context.Background(), User{ID: "Teste"})
	if err != nil {
		panic(err) // panic: cannot marshal type primitive.M to a BSON Document: no encoder found for primitive.M
	}
}
