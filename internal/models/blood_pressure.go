package models

import (
	"time"

	"gorm.io/gorm"
)

type BloodPressure struct {
	gorm.Model
	Sys    uint16    `gorm:"not null"`
	Dia    uint16    `gorm:"not null"`
	UserID uint      `gorm:"not null"`
	Time   time.Time `gorm:"not null"`
}
