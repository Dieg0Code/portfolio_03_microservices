package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	var db *gorm.DB
	var err error

	maxAttempts := 5
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		logrus.WithError(err).Errorf("Failed to connect to database (attempt %d/%d)", attempts, maxAttempts)
		time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
	}

	if err != nil {
		panic("Failed to connect to database after multiple attempts!")
	}

	return db
}

func CheckDatabaseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithError(err).Error("Error getting database connection")
		return err
	}
	return sqlDB.Ping()
}
