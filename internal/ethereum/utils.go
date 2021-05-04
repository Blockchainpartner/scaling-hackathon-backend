package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var walletMutex sync.Mutex

// GetClient initializes a connection to the Ethereum Network
func GetClient() *ethclient.Client {
	nodeURL := utils.EthNodeURI
	client, _ := ethclient.Dial(nodeURL)
	return client
}

func GetWallet() (*bind.TransactOpts, error) {
	walletMutex.Lock()

	pk, err := crypto.HexToECDSA(utils.EthPrivateKey)
	if err != nil {
		return nil, err
	}
	publicKeyECDSA, ok := (pk.Public()).(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	nonce, err := GetClient().PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	gasPrice, err := GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(3))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}

func ReleaseWallet() {
	walletMutex.Unlock()
}
