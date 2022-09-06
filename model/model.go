package model

type Model struct {
	ID        string `json:"id" gorm:"primary_key"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}
