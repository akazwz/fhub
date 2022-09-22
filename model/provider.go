package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Provider struct {
	Model
	ContentHash string `json:"content_hash" gorm:"not null;size:191"`
	Provider    string `json:"provider" gorm:"not null"`
	Key         string `json:"key" gorm:"not null"`
}

func (provider *Provider) TableName() string {
	return "providers"
}

func (provider *Provider) BeforeCreate(*gorm.DB) error {
	id, err := gonanoid.Generate(alphabet, 32)
	provider.ID = id
	return err
}

func (provider *Provider) Create(db *gorm.DB) error {
	return db.Create(provider).Error
}

func (provider *Provider) FindProviderByContentHash(db *gorm.DB) error {
	return db.Where("content_hash = ?", provider.ContentHash).First(provider).Error
}

func FindProvidersByHash(db *gorm.DB, hash string) ([]Provider, error) {
	providers := make([]Provider, 0)
	err := db.Where("content_hash = ?", hash).Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}
