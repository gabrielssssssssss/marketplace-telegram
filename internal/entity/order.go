package entity

import "time"

type Order struct {
	ID          int64
	UserID      int64
	ProductName string
	Amount      float64
	OrderAt     time.Time
}
