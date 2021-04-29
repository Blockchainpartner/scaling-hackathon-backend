package ethereum

import (
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetClient initializes a connection to the Ethereum Network
func GetClient() *ethclient.Client {
	nodeURL := utils.EthNodeURI
	client, _ := ethclient.Dial(nodeURL)
	return client
}
