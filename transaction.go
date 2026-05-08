package billing

import (
	"time"

	"github.com/google/uuid"
)

type (
	Transaction struct {
		Status    string     `json:"status"`
		FromUUID  *uuid.UUID `json:"from_uuid,omitempty"`
		ToUUID    *uuid.UUID `json:"to_uuid,omitempty"`
		Amount    string     `json:"amount"`
		Currency  string     `json:"currency"`
		Timestamp time.Time  `json:"time_stamp"`
	}
)
