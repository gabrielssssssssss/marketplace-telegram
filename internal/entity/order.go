package entity

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID      string
	UserID  int64
	Product json.RawMessage
	Amount  float64
	OrderAt time.Time
}
