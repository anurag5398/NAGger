package entities

import "time"

type BaseEntity struct {
	ID        int64     `gorm:"primary_key" swaggertype:"number" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
