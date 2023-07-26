package utils

import (
	"alfa/db"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	Config := struct {
		Host     string
		Port     int
		User     string
		Password string
	}{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		Config.Host, Config.Port, Config.User, Config.Password)
	dbConn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConn.AutoMigrate(&db.Transaction{}, &db.Advance{})
	return dbConn, nil
}
