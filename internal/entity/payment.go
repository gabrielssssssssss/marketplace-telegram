package entity

import "time"

type Payment struct {
	ID                 *string
	UserID             *int64
	ValueCoin          *float64
	ValueForwardedCoin *float64
	Currency           *string
	Status             *string
	AddressIn          *string
	AddressOut         *string
	TxidIn             *string
	TxidOut            *string
	CreatedAt          time.Time
	ConfirmedAt        time.Time
}

type PaymentCallback struct {
	PaymentID          string
	AddressIn          string
	AdddressOut        string
	Coin               string
	ValueCoin          string
	ValueForwardedCoin string
	TxidIn             string
	TxidOut            string
	Confirmations      string
	Status             string
}
