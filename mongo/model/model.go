package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Collection names
	CollCompany = "company"
	CollUser    = "users"
	CollPets    = "pets"
)

type Entity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// CreatedAt time.Time          `bson:"created_at,omitempty"`
	// UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

type Company struct {
	Entity `bson:"inline"`
	Name   string `bson:"name"`
}

type User struct {
	Entity  `bson:"inline"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
	// CompanyID primitive.ObjectID `bson:"company_id"`
	Company Company `bson:"company"`
	Pets    []*Pet  `bson:"pets,omitempty"`
}

type Pet struct {
	Entity `bson:"inline"`
	Name   string `bson:"name"`
}
