package model

import (
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Capacity struct {
	Model
	UID   string `json:"uid" gorm:"unique"`
	Total int64  `json:"total"`
	Used  int64  `json:"used"`
	Album int64  `json:"album"`
	Video int64  `json:"video"`
}

func (c *Capacity) TableName() string {
	return "capacity"
}

func (c *Capacity) BeforeCreate(*gorm.DB) error {
	id, err := gonanoid.Generate(alphabet, 32)
	c.ID = id
	return err
}

func (c *Capacity) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

func (c *Capacity) FindCapacity(db *gorm.DB) error {
	return db.Where("uid = ?", c.UID).First(c).Error
}

func (c *Capacity) AddUsed(db *gorm.DB, size int64) error {
	err := c.FindCapacity(db)
	if err != nil {
		return err
	}
	if c.Used+size > c.Total {
		return errors.New("capacity not enough")
	}
	return db.Model(&Capacity{}).Where("uid = ?", c.UID).UpdateColumn("used", c.Used+size).Error
}
