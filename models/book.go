package models

import "time"

type Book struct {
	ID          uint      `gorm:"primaryKey" json: "id"`
	Name        string    `json: "name"`
	Description string    `json: "description"`
	MediumPrice float32   `json: "medium_price"`
	Author      string    `json: "author"`
	ImageURL    string    `json: "image_url"`
	CreatedAt   time.Time `json: "created"`
	UpdatedAt   time.Time `json: "updated"`
	DeletedAt   time.Time `gorm: "index" json: "deleted"`
}
