package entity

import "time"

type Payment struct {
	ID          string
	UserID      int64
	Amount      float64
	Currency    string
	TxID        string
	Status      string
	AddressIn   string
	AddressOut  string
	CreatedAt   time.Time
	ConfirmedAt time.Time
}
