package db

import (
	"context"
	"log"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/microgolang/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// InitClient initializes a DB client object
func InitClient() {
	logs.Info(`Init database ...`)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(utils.MongodbURI))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v\n", err)
	}
	// create DB client object
	db = client.Database(utils.MongodbDBName)
	logs.Success(`Database initialized !`)
}
