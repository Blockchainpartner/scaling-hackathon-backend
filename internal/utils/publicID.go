package utils

import (
	"crypto/sha256"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetUUIDFromID is an helper function used to generate an UUID
//based on the ObjectID
func GetUUIDFromID(id primitive.ObjectID) string {
	hash := sha256.Sum256([]byte(id.String()))
	trimmedHash := hash[:16]
	finalUUID, _ := uuid.FromBytes(trimmedHash)
	return finalUUID.String()
}

//GetPublicIDString is an helper function used to generate an UUID based on a string
func GetPublicIDString(str string) string {
	hash := sha256.Sum256([]byte(str))
	trimmedHash := hash[:16]
	finalUUID, _ := uuid.FromBytes(trimmedHash)
	return finalUUID.String()
}
