package model

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID      string          `json:"order_id"`
	UserID  int64           `json:"user_id"`
	Product json.RawMessage `json:"product"`
	Amount  float64         `json:"amount"`
	OrderAt time.Time       `json:"order_at"`
}

type OrderResponse struct {
	Message string `json:"message"`
	Data    Order  `json:"data"`
}

type OrdersResponse struct {
	Message string  `json:"message"`
	Data    []Order `json:"data"`
}
