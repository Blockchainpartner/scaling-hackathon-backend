package models

import (
	"context"
	"errors"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/db"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/microgolang/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRegistry represent the informations about the user's registry keys
type UserRegistry struct {
	Key    *string `json:"key" bson:"key"`
	Secret *string `json:"secret" bson:"secret"`
}

// KYCDob represent the informations about the user's KYC Dob field
type KYCDob struct {
	Date *string `json:"date" bson:"date"`
	Age  uint8   `json:"age" bson:"age"`
}

// KYCLocation represent the informations about the user's KYC Location field
type KYCLocation struct {
	City     *string `json:"city,omitempty" bson:"city"`
	Country  *string `json:"country,omitempty" bson:"country"`
	Postcode *string `json:"postcode,omitempty" bson:"postcode"`
	State    *string `json:"state,omitempty" bson:"state"`
	Street   struct {
		Name   *string `json:"name,omitempty" bson:"name"`
		Number *int    `json:"number,omitempty" bson:"number"`
	} `json:"street" bson:"street"`
}

// KYC represent the informations about the user's KYC
type KYC struct {
	Name     *string     `json:"name" bson:"name"`
	Nat      *string     `json:"nat" bson:"nat"`
	Phone    *string     `json:"phone" bson:"phone"`
	Cell     *string     `json:"cell" bson:"cell"`
	Email    *string     `json:"email" bson:"email"`
	Gender   *string     `json:"gender" bson:"gender"`
	Disabled *bool       `json:"disabled" bson:"disabled"`
	Dob      KYCDob      `json:"dob" bson:"dob"`
	Location KYCLocation `json:"location" bson:"location"`
}

// User is the base of data to create a user
type User struct {
	ID         *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UUID       *string             `json:"UUID,omitempty" bson:"UUID"`
	Password   *string             `json:"password" bson:"password"`
	IsVerified *bool               `json:"isVerified" bson:"isVerified"`
	KYC        KYC                 `json:"KYC" bson:"KYC"`
	Registries []UserRegistry      `json:"registries" bson:"registries"`
}

// NewUser create a new User Object
func NewUser() (x *User) {
	return &User{}
}

//Init will Init the User element to x
func (x *User) Init() *User {
	id := primitive.NewObjectID()

	newElement := &User{
		ID:         &id,
		IsVerified: utils.BoolToPtr(false),
	}
	return newElement
}

//Post will save the User element in the database
func (x *User) Post() error {
	var err error

	if db.Users != nil {
		_, err = db.Users.InsertOne(context.Background(), x)
	} else {
		return errors.New(`database not initialized`)
	}
	return err
}

//FindByUUID will perform a search in the Users collection based on the UUID
func (x *User) FindByUUID(UUID string) (*User, error) {
	document := db.Users.FindOne(
		context.Background(),
		bson.M{`UUID`: UUID},
	)
	element := User{}
	err := document.Decode(&element)
	return &element, err
}

//FindBy will perform a search in the Users collection based on by
func (x *User) FindBy(by bson.M) (*User, error) {
	document := db.Users.FindOne(context.Background(), by)
	element := User{}
	err := document.Decode(&element)
	return &element, err
}

//List will perform a search in the Users collection to find all the elements
func (x *User) List() ([]User, error) {
	query, err := db.Users.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func() {
		err := query.Close(context.Background())
		//Nothing can be done by the process on error
		if err != nil {
			logs.Error(err.Error())
		}
	}()

	elements := make([]User, 0)
	if err := query.All(context.Background(), &elements); err != nil {
		return nil, err
	}
	return elements, query.Err()
}

//ListBy will perform a search in the Users collection to find all the elements
func (x *User) ListBy(by bson.M) ([]User, error) {
	query, err := db.Users.Find(context.Background(), by)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := query.Close(context.Background())
		//Nothing can be done by the process on error
		if err != nil {
			logs.Error(err.Error())
		}
	}()

	elements := make([]User, 0)
	if err := query.All(context.Background(), &elements); err != nil {
		return nil, err
	}
	return elements, query.Err()
}

//SelfUpdate will perform an update on one User based on X
func (x *User) SelfUpdate() error {
	var err error

	if db.Users != nil {
		_, err = db.Users.UpdateOne(
			context.Background(),
			bson.M{`UUID`: x.UUID},
			bson.M{`$set`: x},
		)
	} else {
		return errors.New(`database not initialized`)
	}
	return err
}

//Exists will check if a specific User exists
func (x *User) Exists(conditions bson.M) bool {
	count, err := db.Users.CountDocuments(context.Background(), conditions)
	if err != nil {
		return false
	}
	return count > 0
}

//Count will count the number of element matching conditions
func (x *User) Count(conditions bson.M) (int64, error) {
	count, err := db.Users.CountDocuments(context.Background(), conditions)
	return count, err
}
