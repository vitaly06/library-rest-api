package entities

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at" gorm:"AutoCreatedAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"AutoUpdatedAt"`
}
