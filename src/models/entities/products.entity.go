package entities

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Code      string         `gorm:"column:code" json:"code"`
	Price     uint           `gorm:"column:price" json:"price"`
}
