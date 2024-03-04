package models

import (
	"time"

	"gorm.io/gorm"
)

type WaterIntake struct {
	gorm.Model
	Cups   float32   `gorm:"not null"`
	UserID uint      `gorm:"not null"`
	Time   time.Time `gorm:"not null"`
}
