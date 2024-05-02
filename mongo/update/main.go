package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   any    `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:MongoPass321!@localhost:27017"))
	if err != nil {
		panic(err)
	}

	updateId, _ := primitive.ObjectIDFromHex("6626b9e37f7f187122b6c920")
	user := User{
		ID:   updateId,
		Name: "Teste 1234",
	}

	filter := bson.M{"_id": updateId}
	update := bson.M{"$set": user}

	coll := client.Database("test").Collection("coll")
	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}
