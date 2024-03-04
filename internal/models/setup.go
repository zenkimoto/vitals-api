package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase(host string, user string, password string, dbname string) {
	// TODO: Add support for different database types.
	const databaseType = "postgres"

	var database *gorm.DB
	var err error

	if databaseType == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbname)
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if databaseType == "sqlite" {
		database, err = gorm.Open(sqlite.Open("./vitals.db"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database.")
	}

	err = database.AutoMigrate(&BloodPressure{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&Weight{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&WaterIntake{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&SugarIntake{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return
	}

	DB = database
}
