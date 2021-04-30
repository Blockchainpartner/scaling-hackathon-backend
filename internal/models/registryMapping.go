package models

import (
	"context"
	"errors"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/db"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/microgolang/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RegistryMapping is the base of data to create a registry
type RegistryMapping struct {
	ID            *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UUID          *string             `json:"UUID" bson:"UUID"`
	RegistryKey   *string             `json:"registryKey" bson:"registryKey"`
	Identity      *string             `json:"identity" bson:"identity"`
	IdentityIndex *uint64             `json:"identityIndex" bson:"identityIndex"`
}

// NewRegistryMapping create a new RegistryMapping Object
func NewRegistryMapping() (x *RegistryMapping) {
	return &RegistryMapping{}
}

//Init will Init the RegistryMapping element to x
func (x *RegistryMapping) Init() *RegistryMapping {
	id := primitive.NewObjectID()
	UUID := utils.GetUUIDFromID(id)

	newElement := &RegistryMapping{
		ID:   &id,
		UUID: utils.StrToPtr(UUID),
	}
	return newElement
}

//Post will save the RegistryMapping element in the database
func (x *RegistryMapping) Post() error {
	var err error

	if db.RegistriesMapping != nil {
		_, err = db.RegistriesMapping.InsertOne(context.Background(), x)
	} else {
		return errors.New(`database not initialized`)
	}
	return err
}

//FindByUUID will perform a search in the Users collection based on the UUID
func (x *RegistryMapping) FindByUUID(UUID string) (*RegistryMapping, error) {
	document := db.RegistriesMapping.FindOne(
		context.Background(),
		bson.M{`UUID`: UUID},
	)
	element := RegistryMapping{}
	err := document.Decode(&element)
	return &element, err
}

//FindBy will perform a search in the Users collection based on by
func (x *RegistryMapping) FindBy(by bson.M) (*RegistryMapping, error) {
	document := db.RegistriesMapping.FindOne(context.Background(), by)
	element := RegistryMapping{}
	err := document.Decode(&element)
	return &element, err
}

//List will perform a search in the Users collection to find all the elements
func (x *RegistryMapping) List() ([]RegistryMapping, error) {
	query, err := db.RegistriesMapping.Find(context.Background(), bson.M{})
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

	elements := make([]RegistryMapping, 0)
	if err := query.All(context.Background(), &elements); err != nil {
		return nil, err
	}
	return elements, query.Err()
}

//ListBy will perform a search in the Users collection to find all the elements
func (x *RegistryMapping) ListBy(by bson.M) ([]RegistryMapping, error) {
	query, err := db.RegistriesMapping.Find(context.Background(), by, options.Find().SetSort(bson.M{`identityIndex`: -1}))
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

	elements := make([]RegistryMapping, 0)
	if err := query.All(context.Background(), &elements); err != nil {
		return nil, err
	}
	return elements, query.Err()
}

//SelfUpdate will perform an update on one RegistryMapping based on X
func (x *RegistryMapping) SelfUpdate() error {
	var err error

	if db.RegistriesMapping != nil {
		_, err = db.RegistriesMapping.UpdateOne(
			context.Background(),
			bson.M{`UUID`: x.UUID},
			bson.M{`$set`: x},
		)
	} else {
		return errors.New(`database not initialized`)
	}
	return err
}

//Exists will check if a specific RegistryMapping exists
func (x *RegistryMapping) Exists(conditions bson.M) bool {
	count, err := db.RegistriesMapping.CountDocuments(context.Background(), conditions)
	if err != nil {
		return false
	}
	return count > 0
}

//Count will count the number of element matching conditions
func (x *RegistryMapping) Count(conditions bson.M) (int64, error) {
	count, err := db.RegistriesMapping.CountDocuments(context.Background(), conditions)
	return count, err
}
