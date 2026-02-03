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
