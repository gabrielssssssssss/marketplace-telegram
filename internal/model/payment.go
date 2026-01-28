package model

import "time"

type Payment struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	TxID        string    `json:"tx_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	ConfirmedAt time.Time `json:"confirmed_at"`
}
