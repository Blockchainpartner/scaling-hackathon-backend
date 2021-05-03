package ethereum

import (
	"context"
	"errors"
	"time"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/models"
	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/gin-gonic/gin"
	"github.com/microgolang/logs"
)

var errNotFound = errors.New(`not found`)
var errTimeout = errors.New(`timeout`)

//TxSubsOptions is the structure containing the options used while listening for
//a transaction completion.
//- Timeout is the maximum time during which the process will try to get an answer
//- Delay is the amount of time the process will sleep before trying to get another
//  answer (in the same try)
//- MaxRetries is the maximum amount of retry we authorize to get a valid completion
//  answer from the blockchain. 1 try duration == Timeout variable.
//- CurrentRetry is an incremental counter, indicating the current retry process.
type TxSubsOptions struct {
	Timeout      time.Duration
	Delay        time.Duration
	MaxRetries   int
	CurrentRetry int
}

func setTxOptions() TxSubsOptions {
	return TxSubsOptions{
		Timeout:      60 * time.Second,
		Delay:        4 * time.Second,
		MaxRetries:   54,
		CurrentRetry: 0,
	}
}

func retreiveReceipt(txHash common.Hash) (*types.Receipt, error) {
	receipt, err := GetClient().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
func getTransactionStatus(txHash common.Hash, opts TxSubsOptions) (*types.Receipt, *types.Transaction, bool, error) {
	var isPending bool
	var err error
	var receipt *types.Receipt
	var transaction *types.Transaction

	timeout := time.Now().Add(opts.Timeout)

	for time.Now().Before(timeout) {
		/**********************************************************************
		**	We first need to check if the transaction exists in the blockchain,
		**	aka if it is pending, but existing. On error, the error will no
		**	longer be err, but 'not found'
		**********************************************************************/
		transaction, isPending, err = GetClient().TransactionByHash(context.Background(), txHash)
		if isPending {
			time.Sleep(opts.Delay)
			continue
		} else if err != nil && !isPending {
			if err == errNotFound {
				return nil, transaction, false, errNotFound
			}
			return nil, transaction, false, err
		}

		/**********************************************************************
		**	Once the transaction is no longer pending, we can get the receipt
		**	and take action from it
		**********************************************************************/
		receipt, err = retreiveReceipt(txHash)
		if err != nil {
			return receipt, transaction, false, err
		}
		return receipt, transaction, false, nil
	}
	return nil, transaction, isPending, errTimeout
}
func subscribePendingTransactions(txID *string, retry uint8) (*types.Receipt, *types.Transaction, bool, error) {
	tx := common.HexToHash(*txID)
	opts := setTxOptions()
	/**************************************************************************
	**	This function will try, multiple time, to get the result of a
	**	transaction.
	**	- On success, the transaction receipt is returned.
	**	- On timeout (transaction still pending), we retry the subscription
	**	  opts.MaxRetries times for opts.Timeout seconds.
	**	- On MaxRetries reached, a nil receipt and an error is returned.
	**	- On error, a nil receipt and the error is returned.
	**
	**	Special cases :
	**	- We got another transaction, with the same Nonce, in the DB. If the
	**	  TxTime of the new one is after the one of the old one, we should stop
	**	  listening for the old one.
	**	- We got another transaction, with a bigger Nonce,
	**************************************************************************/
	transactionReceipt, transaction, isPending, transactionError := getTransactionStatus(tx, opts)

	if isPending {
		logs.Info(`transaction not mined, still pending`)
		/******************************************************************
		**	We didn't get the transaction because of the function timeout,
		**	and the transaction is still pending
		**	While the MaxRetries variable is not reached, we will pause
		**	the process for some seconds, and try to get the transaction
		** result after that
		******************************************************************/
		if opts.CurrentRetry < opts.MaxRetries {
			logs.Info(`Sleep until next retry`)
			//If the opt speed is normal, we will wait for 30 x the currentRetry, which
			//can lead to a maximum listening time of 44584 sec
			delayBeforeNextLoop := time.Duration(opts.CurrentRetry) * 30 * time.Second
			time.Sleep(delayBeforeNextLoop)
			opts.CurrentRetry++
			logs.Info(`Retrying ...`)

			return subscribePendingTransactions(txID, retry+1)
		}

		/******************************************************************
		**	If we reach this point, the transaction is still pending after
		**	the max retries (21h for slow, 10mn for medium and 2.4mn for
		** fast). We should proceed another way.
		******************************************************************/
		logs.Warning(`the transaction is still pending and we reached a timeout`)
		return transactionReceipt, transaction, true, transactionError
	} else if transactionError != nil {
		/******************************************************************
		**	If we reach this point, the transaction has failed for some
		**	reason. The main thread should decide what to do with it
		******************************************************************/
		if transactionError.Error() == `not found` && retry < 10 {
			time.Sleep(time.Second * 5)
			return subscribePendingTransactions(txID, retry+1)
		}
		return transactionReceipt, transaction, false, transactionError
	}

	/******************************************************************
	**	If we reach this point, the transaction has been mined without
	**	error (but may not be with a 1 status)
	******************************************************************/
	return transactionReceipt, transaction, false, nil
}

//WatchTransaction will take a specific transaction model and watch for the blockchain updates
//to get some information about this transaction (still pending, success, failed, error)
func WatchTransaction(tx *models.Transaction) bool {
	logs.Info(`Watching transaction : ` + *tx.TxID)
	receipt, _, isPending, err := subscribePendingTransactions(tx.TxID, 0)

	var isFailed = types.ReceiptStatusFailed
	var isSucceed = types.ReceiptStatusSuccessful

	if isPending {
		/**********************************************************************
		**	Transaction has not be mined yet
		**********************************************************************/
		logs.Warning(`transaction is still pending`)
		return false
	}
	if err != nil {
		if err == vm.ErrOutOfGas {
			/******************************************************************
			**	The transaction has failed because it is laking some gas.
			**	We should retry it with an hight gas
			******************************************************************/
			logs.Error(err.Error())
			return false
		}
		/**********************************************************************
		**	Error with the transaction, we log it in the database and move
		**	to the next transaction.
		**********************************************************************/
		logs.Error(err.Error())
		tx.IsPending = &isPending
		tx.IsInQueue = utils.BoolToPtr(false)
		tx.Status = &isFailed
		return false
	} else if receipt.Status == isFailed {
		/**********************************************************************
		**	The transaction has been successfully mined, but the status is
		**	failed. We log it in the database and move to the next
		**	transaction
		**	We also increment the wallet nonce.
		**********************************************************************/
		logs.Error(errors.New(`transaction mined, but not successful: ` + *tx.TxID))
		utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
			`registry`: tx.Key,
			`step`:     `transaction mined, but not successful`,
			`type`:     `error`,
		})
		tx.IsPending = &isPending
		tx.IsInQueue = utils.BoolToPtr(false)
		tx.Status = &isFailed
		return false
	}
	/**************************************************************************
	**	The transaction has been successfully mined, and is successful !
	**	Depending on the type of transaction, we can execute the needed
	**	actions.
	**	We log this information in the DB and move to the next transaction.
	**	We also increment the wallet nonce.
	**************************************************************************/
	logs.Info(`transaction success`)
	tx.IsPending = &isPending
	tx.IsInQueue = utils.BoolToPtr(false)
	tx.Status = &isSucceed
	utils.NewPusher().Identities.Push(`PROCESS`, gin.H{
		`registry`: tx.Key,
		`step`:     `Transaction successful !`,
		`type`:     `success`,
	})
	return true
}
