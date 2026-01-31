package model

import "time"

type Payment struct {
	ID                 string     `json:"id"`
	UserID             int64      `json:"user_id"`
	ValueCoin          float64    `json:"value_coin"`
	ValueForwardedCoin float64    `json:"value_forwarded_coin"`
	Currency           string     `json:"currency"`
	Status             string     `json:"status"`
	AddressIn          string     `json:"address_in"`
	AddressOut         string     `json:"address_out"`
	TxidIn             string     `json:"txid_in"`
	TxidOut            string     `json:"txid_out"`
	CreatedAt          *time.Time `json:"created_at"`
	ConfirmedAt        *time.Time `json:"confirmed_at"`
}
