package billing

import (
	"time"

	"github.com/google/uuid"
)

type (
	Billing interface {
		GetBalance() int64
		GetTransactions() []*Transaction
	}

	billing struct {
		UUID      uuid.UUID
		balance   int64
		timestamp time.Time
	}
)

func NewBilling() Billing {
	billing := &billing{}

	return billing
}

func (billing *billing) GetBalance() int64 {
	return billing.balance
}

func (billing *billing) GetTransactions() []*Transaction {
	var result []*Transaction

	return result
}
