package model

import (
	"time"
)

type Order struct {
	ID      int64          `json:"id"`
	UserID  int64          `json:"user_id"`
	Product map[string]any `json:"product"`
	Amount  float64        `json:"amount"`
	OrderAt time.Time      `json:"order_at"`
}
