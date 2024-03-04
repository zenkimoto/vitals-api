package models

import (
	"time"

	"gorm.io/gorm"
)

type Weight struct {
	gorm.Model
	Weight float32   `gorm:"not null"`
	UserID uint      `gorm:"not null"`
	Time   time.Time `gorm:"not null"`
}
