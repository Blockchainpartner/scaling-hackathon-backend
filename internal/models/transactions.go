package models

import (
	"time"

	"github.com/Blockchainpartner/scaling-hackathon-backend/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Transaction is the model used to represent an Ethereum transaction
type Transaction struct {
	ID        *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedAt *time.Time          `json:"-" bson:"createdAt"`
	UpdatedAt *time.Time          `json:"-" bson:"updatedAt"`
	From      *string             `json:"from" bson:"from"`
	TxID      *string             `json:"txID" bson:"txID"`
	Key       *string             `json:"key" bson:"key"`
	TxTime    time.Time           `json:"txTime" bson:"txTime"`
	Status    *uint64             `json:"status" bson:"status"`
	IsPending *bool               `json:"isPending" bson:"isPending"`
	IsInQueue *bool               `json:"isInQueue" bson:"isInQueue"`
}

//TransactionFilters is the model used to query the transactions in the database
type TransactionFilters struct {
	ID        *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	From      *string             `json:"from,omitempty" bson:"from,omitempty"`
	TxID      *string             `json:"txID,omitempty" bson:"txID,omitempty"`
	Key       *string             `json:"key,omitempty" bson:"key,omitempty"`
	TxTime    time.Time           `json:"txTime,omitempty" bson:"txTime,omitempty"`
	Status    *uint64             `json:"status,omitempty" bson:"status,omitempty"`
	IsPending *bool               `json:"isPending,omitempty" bson:"isPending,omitempty"`
	IsInQueue *bool               `json:"isInQueue,omitempty" bson:"isInQueue,omitempty"`
}

// NewTransaction create a new Transaction Object
func NewTransaction() (x *Transaction) {
	return &Transaction{}
}

//Init will Init the TransactionFilter element to x
func (x *Transaction) Init() *Transaction {
	id := primitive.NewObjectID()
	newElement := &Transaction{
		ID:        &id,
		CreatedAt: utils.TimeToPtr(time.Now()),
		UpdatedAt: utils.TimeToPtr(time.Now()),
		IsPending: utils.BoolToPtr(false),
		IsInQueue: utils.BoolToPtr(false),
	}
	return newElement
}
