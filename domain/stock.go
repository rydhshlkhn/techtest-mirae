package domain

import "time"

type Stock struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Price     float64   `json:"price"`
	Frequency int       `json:"frequency"`
	Volume    int       `json:"volume"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
