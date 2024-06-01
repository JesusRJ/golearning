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

type AbstractEntity interface {
	GetID() primitive.ObjectID
}

type Entity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// CreatedAt time.Time          `bson:"created_at,omitempty"`
	// UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func (e *Entity) GetID() primitive.ObjectID {
	return e.ID
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
	// Tag couldn't be "company_id" because of the ref field name
	Company *Company `bson:"company" ref:"belongsTo,company,company_id,_id,companies"`
	Pets    []*Pet   `bson:"pets" ref:"hasMany,pet,_id,user_id,pets"`
}

type Address struct {
	Street string `bson:"street"`
	Number int    `bson:"number"`
}

type Phone struct {
	Number string `bson:"number"`
}

type Pet struct {
	Entity `bson:"inline"`
	User   *User  `bson:"user" ref:"belongsTo,user,user_id,_id,user"`
	Name   string `bson:"name"`
}

func (c Company) String() string { return c.Name }

func (a Address) String() string { return fmt.Sprintf("%v (%v)", a.Street, a.Number) }

func (p Phone) String() string { return p.Number }

func (p Pet) String() string { return p.Name }
