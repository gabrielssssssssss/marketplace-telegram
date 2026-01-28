package entity

import "time"

type Payment struct {
	ID          int64
	UserID      int64
	Amount      float64
	Currency    string
	TxID        string
	Status      string
	CreatedAt   time.Time
	ConfirmedAt time.Time
}
