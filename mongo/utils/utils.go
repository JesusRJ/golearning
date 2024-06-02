package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCompanyID() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6632bff5465065406a8f320a")
	return id
}
