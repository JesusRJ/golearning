package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Collection names
	CollCompany = "company"
	CollUser    = "user"
	CollPets    = "pet"
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
	Name    string   `bson:"name"`
	Address *Address `bson:"address"` // embedded
	Phone   []*Phone `bson:"phones"`  // embedded
	Company *Company `bson:"company"`
	Pets    []*Pet   `bson:"pets"`
}

type Address struct {
	Street string `bson:"street"`
	Number int    `bson:"number"`
}

type Phone struct {
	User   *User  `bson:"-"`
	Number string `bson:"number"`
}

type Pet struct {
	Entity `bson:"inline"`
	Name   string `bson:"name"`
}

func (c Company) String() string {
	return c.Name
}

func (a Address) String() string {
	return fmt.Sprintf("%v (%v)", a.Street, a.Number)
}

func (p Phone) String() string {
	return p.Number
}

func (p Pet) String() string {
	return p.Name
}
