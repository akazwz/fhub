package model

import "time"

type Model struct {
	ID        string    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"UpdatedAt"`
}
