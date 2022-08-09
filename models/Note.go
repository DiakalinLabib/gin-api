package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Content   string         `json:"content"`
	IsDone    bool           `json:"is_done"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
