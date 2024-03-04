package models

import (
	"time"

	"gorm.io/gorm"
)

type SugarIntake struct {
	gorm.Model
	Grams  uint      `gorm:"not null"`
	UserID uint      `gorm:"not null"`
	Time   time.Time `gorm:"not null"`
}
