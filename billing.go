package billing

import (
	"time"

	"github.com/google/uuid"
)

type (
	Billing interface {
		GetBalance() int64
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
