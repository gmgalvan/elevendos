package entity

import (
	"time"
)

type Product struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Comments  *string    `json:"comments"`
	Timestamp *time.Time `json:"timestamp"`
}
