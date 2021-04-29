package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	//Users is the collection used to store the Users
	Users *mongo.Collection
	//Registries is the collection used to store the Registries
	Registries *mongo.Collection
	//RegistriesMapping is the collection used to store the relation user <-> registry
	RegistriesMapping *mongo.Collection
)

// InitCollections initializes collection objects and creates indexes
func InitCollections() {
	Users = db.Collection("users")
	Registries = db.Collection("registries")
	RegistriesMapping = db.Collection("registries_mapping")
}
