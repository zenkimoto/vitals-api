package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName         string
	LastName          string
	Role              string
	UserName          string `gorm:"uniqueIndex,not null"`
	PasswordHash      string
	BloodPressureList []BloodPressure `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	WeightList        []Weight        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	WaterIntakeList   []WaterIntake   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SugarIntakeList   []SugarIntake   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
