package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCompanyID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")
	return id
}

func GetUserID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("661f17bffc35c18b2f85e975")
	return id
}
