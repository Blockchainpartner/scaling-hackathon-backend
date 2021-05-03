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
	// EthNodeURIHttp : Uri of the ethereum node in http mode
	EthNodeURIHttp string
	// EthPrivateKey : PrivateKey used for this demonstration
	EthPrivateKey string
	// EthContractAddress : address of our smartcontract
	EthContractAddress string

	// PusherSecret represent the pusher secret
	PusherSecret string
	// PusherKey represent the pusher key
	PusherKey string
	// PusherID represent the pusher ID
	PusherID string
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
	EthNodeURIHttp, exists = os.LookupEnv("ETH_NODE_URI_HTTP")
	if !exists {
		log.Fatal("ETH_NODE_URI_HTTP environment variable not set")
	}
	EthPrivateKey, exists = os.LookupEnv("ETH_PRIVATE_KEY")
	if !exists {
		log.Fatal("ETH_PRIVATE_KEY environment variable not set")
	}
	EthContractAddress, exists = os.LookupEnv("ETH_CONTRACT_ADDRESS")
	if !exists {
		log.Fatal("ETH_CONTRACT_ADDRESS environment variable not set")
	}
}

func initAPIKeys() {
	var exists bool

	PusherSecret, exists = os.LookupEnv("PUSHER_SECRET")
	if !exists {
		log.Fatal("PUSHER_SECRET environment variable not set")
	}
	PusherKey, exists = os.LookupEnv("PUSHER_KEY")
	if !exists {
		log.Fatal("PUSHER_KEY environment variable not set")
	}
	PusherID, exists = os.LookupEnv("PUSHER_ID")
	if !exists {
		log.Fatal("PUSHER_ID environment variable not set")
	}
}

// InitEnvironment initializes variables from environment
func InitEnvironment() {
	initMongoDB()
	initEthCore()
	initAPIKeys()
}
