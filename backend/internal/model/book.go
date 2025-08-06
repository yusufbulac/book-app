package model

import (
	"time"
)

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Author    string    `gorm:"type:varchar(255);not null" json:"author"`
	Year      int       `gorm:"not null" json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
