package env

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/zenkimoto/vitals-server-api/internal/util"
)

// Load environment variables from .env file if it exists
// Mainly for development purposes only.
func LoadEnvironmentVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Print("No .env file found. Using environment variables...")
	} else {
		log.Print("Environment variables loaded from .env file.")
	}
}

// Port Section

// Get port defined in the PORT environment variable.
// If the PORT environment variable does not exist,
// default to "localhost:8080"
func GetPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		return "localhost:8080"
	} else {
		return port
	}
}

// Database Section

// Get database connection information from environment variables
func GetDatabaseConnectionInfo() (string, string, string, string) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return host, user, password, dbname
}

// JWT Key Section

var jwtKey string = ""

func GetJWTKey() string {
	if jwtKey != "" {
		return jwtKey
	}

	jwtKey = os.Getenv("JWT_KEY")

	if jwtKey == "" {
		log.Print("ERROR: Unable to retrieve environment variable JWT_KEY")
		log.Print("Generating random key...")

		jwtKey = util.RandString(32)
	}

	return jwtKey
}

// Token Expiration Duration Section

var duration time.Duration

func GetTokenExpirationDuration() time.Duration {
	if duration != 0 {
		return duration
	}

	duration, err := strconv.Atoi(os.Getenv("DURATION_SEC"))

	// Return default duration of 6 hours if unable to get
	// duration from environment variables
	if err != nil {
		log.Print("ERROR: Unable to retrieve environment variable DURATION_SEC")
		log.Print("Setting default token expiration duration of 6 hours")

		return time.Hour * 6
	}

	return time.Duration(duration)
}
