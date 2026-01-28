package model

import "time"

type Orders struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	ProductName string    `json:"product_name"`
	Amount      float64   `json:"amount"`
	OrderAt     time.Time `json:"order_at"`
}
