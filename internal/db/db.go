package db

import (
	"log"

	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/env"
	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ORM *gorm.DB

type Db struct {
	db *gorm.DB
}

func NewDBWithEnv() *gorm.DB {
	connectDBViaEnv()
	return ORM
}

func NewDBWithConfig(conn string) *gorm.DB {
	connectViaConfigFile(conn)
	return ORM
}

func connectViaConfigFile(conn string) {
	logger := logger.NewLogger()

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		logger.ErrorLogger.Printf("Error connecting to Postgres with config file : %+v", err)
		log.Fatalf("DB Connection error with config file: %+v", err)
	}
	logger.InfoLogger.Println("Successfully connected to Database")
	ORM = db
}

func connectDBViaEnv() {

	logger := logger.NewLogger()
	db, err := gorm.Open("postgres", env.GetDBEnv())
	if err != nil {
		logger.ErrorLogger.Printf("Couldn't connect to Database %+v", err)
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	logger.InfoLogger.Println("Successfully connected to Database")

	ORM = db
}
