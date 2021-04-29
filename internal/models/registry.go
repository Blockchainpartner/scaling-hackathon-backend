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

// Registry is the base of data to create a registry
type Registry struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UUID        *string             `json:"UUID" bson:"UUID"`
	Key         *string             `json:"key" bson:"key"`
	Hash        *string             `json:"hash" bson:"hash"`
	Description *string             `json:"description" bson:"description"`
}

// NewRegistry create a new Registry Object
func NewRegistry() (x *Registry) {
	return &Registry{}
}

//Init will Init the Registry element to x
func (x *Registry) Init() *Registry {
	id := primitive.NewObjectID()
	UUID := utils.GetUUIDFromID(id)

	newElement := &Registry{
		ID:   &id,
		UUID: utils.StrToPtr(UUID),
	}
	return newElement
}

//Post will save the Registry element in the database
func (x *Registry) Post() error {
	var err error

	if db.Registries != nil {
		_, err = db.Registries.InsertOne(context.Background(), x)
	} else {
		return errors.New(`database not initialized`)
	}
	return err
}

//FindByUUID will perform a search in the Users collection based on the UUID
func (x *Registry) FindByUUID(UUID string) (*Registry, error) {
	document := db.Registries.FindOne(
		context.Background(),
		bson.M{`UUID`: UUID},
	)
	element := Registry{}
	err := document.Decode(&element)
	return &element, err
}

//FindBy will perform a search in the Users collection based on by
func (x *Registry) FindBy(by bson.M) (*Registry, error) {
	document := db.Registries.FindOne(context.Background(), by)
	element := Registry{}
	err := document.Decode(&element)
	return &element, err
}

//List will perform a search in the Users collection to find all the elements
func (x *Registry) List() ([]Registry, error) {
	query, err := db.Registries.Find(context.Background(), bson.M{})
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

	elements := make([]Registry, 0)
	if err := query.All(context.Background(), &elements); err != nil {
		return nil, err
	}
	return elements, query.Err()
}

//ListBy will perform a search in the Users collection to find all the elements
func (x *Registry) ListBy(by bson.M) ([]Registry, error) {
	query, err := db.Registries.Find(context.Background(), by)
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

	elements := make([]Registry, 0)
	if err := query.All(context.Background(), &elements); err != nil {
		return nil, err
	}
	return elements, query.Err()
}

//SelfUpdate will perform an update on one Registry based on X
func (x *Registry) SelfUpdate() error {
	var err error

	if db.Registries != nil {
		_, err = db.Registries.UpdateOne(
			context.Background(),
			bson.M{`UUID`: x.UUID},
			bson.M{`$set`: x},
		)
	} else {
		return errors.New(`database not initialized`)
	}
	return err
}

//Exists will check if a specific Registry exists
func (x *Registry) Exists(conditions bson.M) bool {
	count, err := db.Registries.CountDocuments(context.Background(), conditions)
	if err != nil {
		return false
	}
	return count > 0
}

//Count will count the number of element matching conditions
func (x *Registry) Count(conditions bson.M) (int64, error) {
	count, err := db.Registries.CountDocuments(context.Background(), conditions)
	return count, err
}
