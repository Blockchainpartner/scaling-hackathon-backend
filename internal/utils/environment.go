package utils

import (
	"log"
	"os"
)

var (
	// MongodbURI : URI of the MongoDB
	MongodbURI string
	// MongodbDBName : name of the DB
	MongodbDBName string

	// EthNodeURI : Uri of the ethereum node
	EthNodeURI string

	// EtherscanAPI represent the APIKey for etherscan
	EtherscanAPI string
)

func initMongoDB() {
	var exists bool
	MongodbURI, exists = os.LookupEnv("MONGODB_URI")
	if !exists {
		log.Fatal("MONGODB_URI environment variable not set")
	}
	MongodbDBName, exists = os.LookupEnv("MONGODB_DB_NAME")
	if !exists {
		log.Fatal("MONGODB_DB_NAME environment variable not set")
	}
}

func initEthCore() {
	var exists bool

	EthNodeURI, exists = os.LookupEnv("ETH_NODE_URI")
	if !exists {
		log.Fatal("ETH_NODE_URI environment variable not set")
	}
}

func initAPIKeys() {
	var exists bool

	EtherscanAPI, exists = os.LookupEnv("API_ETHERSCAN")
	if !exists {
		log.Fatal("API_ETHERSCAN environment variable not set")
	}
}

// InitEnvironment initializes variables from environment
func InitEnvironment() {
	initMongoDB()
	initEthCore()
	// initAPIKeys()
}
