package entities

import "time"

type Record struct {
	ID        int64     `gorm:"primarykey" swaggertype:"number" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}
