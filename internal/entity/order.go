package entity

import "time"

type Order struct {
	ID      *int64
	UserID  *int64
	Details *map[string]any
	Amount  *float64
	OrderAt *time.Time
}
