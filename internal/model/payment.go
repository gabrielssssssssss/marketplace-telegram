package model

import "time"

type Payment struct {
	ID          string    `json:"id"`
	UserID      int64     `json:"user_id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	TxID        string    `json:"tx_id"`
	Status      string    `json:"status"`
	AddressIn   string    `json:"address_in"`
	AddressOut  string    `json:"address_out"`
	CreatedAt   time.Time `json:"created_at"`
	ConfirmedAt time.Time `json:"confirmed_at"`
}
