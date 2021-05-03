/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Monday March 22nd 2021
**	@Filename:				SignClaim.go
******************************************************************************/

package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/contracts"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/microgolang/logs"
)

// UpdateRegistry will update a specific registry with some new identities
func UpdateRegistry(registryKey string, oldRegistryHash string, newRegistryHash string) error {
	pk, err := crypto.HexToECDSA(utils.EthPrivateKey)
	if err != nil {
		logs.Error(err)
		return err
	}
	publicKeyECDSA, ok := (pk.Public()).(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	nonce, err := GetClient().PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		logs.Error(err)
		return err
	}
	gasPrice, err := GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		logs.Error(err)
		return err
	}
	auth := bind.NewKeyedTransactor(pk)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	contract, err := contracts.NewCairoProver(common.HexToAddress(utils.EthContractAddress), GetClient())
	if err != nil {
		logs.Error(err)
		return err
	}

	bigRegistryKey, success := big.NewInt(0).SetString(registryKey, 10)
	if !success {
		return errors.New(`invalid big int`)
	}
	bigOldRegistryHash, success := big.NewInt(0).SetString(oldRegistryHash, 10)
	if !success {
		return errors.New(`invalid big int`)
	}
	bigNewRegistryHash, success := big.NewInt(0).SetString(newRegistryHash, 10)
	if !success {
		return errors.New(`invalid big int`)
	}

	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{`registry`: registryKey, `step`: `Performing update tx ...`, `type`: `info`})

	tx, err := contract.UpdateRegistry(auth, bigRegistryKey, bigOldRegistryHash, bigNewRegistryHash)
	if err != nil {
		logs.Error(err)
		return err
	}

	txID := tx.Hash().Hex()
	transaction := models.NewTransaction().Init()
	transaction.From = utils.StrToPtr(address)
	transaction.TxID = &txID
	transaction.Key = utils.StrToPtr(registryKey)

	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: registryKey,
		`step`:     `Waiting for the transaction to be mined`,
		`type`:     `info`,
	})

	if WatchTransaction(transaction) {
		return nil
	}
	return errors.New(`impossible to watch transaction`)
}
