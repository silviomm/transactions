package database

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"pismo-challenge/models/account"
	"pismo-challenge/models/transaction"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)
	connectionString := os.Getenv("CONNECTION_STRING")
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panic("Error connecting to database", err)
	}
	err = DB.AutoMigrate(&account.Account{})
	if err != nil {
		log.Panic("Error migrating Accounts table", err)
	}
	err = DB.AutoMigrate(&transaction.Transaction{})
	if err != nil {
		log.Panic("Error migrating Transactions table", err)
	}
}
