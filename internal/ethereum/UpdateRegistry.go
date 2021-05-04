/******************************************************************************
**	@Author:				Thomas Bouder <Tbouder>
**	@Email:					Tbouder@protonmail.com
**	@Date:					Monday March 22nd 2021
**	@Filename:				SignClaim.go
******************************************************************************/

package ethereum

import (
	"errors"
	"math/big"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/contracts"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/microgolang/logs"
)

// UpdateRegistry will update a specific registry with some new identities
func UpdateRegistry(registryKey string, oldRegistryHash, newRegistryHash *big.Int) error {
	auth, err := GetWallet()
	if err != nil {
		logs.Error(err)
		return err
	}
	defer ReleaseWallet()

	contract, err := contracts.NewCairoProver(common.HexToAddress(utils.EthContractAddress), GetClient())
	if err != nil {
		logs.Error(err)
		return err
	}

	bigRegistryKey, success := big.NewInt(0).SetString(registryKey, 10)
	if !success {
		return errors.New(`invalid big int`)
	}

	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{`registry`: registryKey, `step`: `Performing update tx ...`, `type`: `info`})

	tx, err := contract.UpdateRegistry(auth, bigRegistryKey, oldRegistryHash, newRegistryHash)
	if err != nil {
		logs.Error(err)
		return err
	}

	txID := tx.Hash().Hex()
	transaction := models.NewTransaction().Init()
	transaction.From = utils.StrToPtr(auth.From.Hex())
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
