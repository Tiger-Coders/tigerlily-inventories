package db

import (
	"log"

	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/env"
	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/injection"
	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ORM *gorm.DB

type Db struct {
	db *gorm.DB
}

func NewDB() *gorm.DB {
	connectDB()
	return ORM
}

func connectDB() {
	config := injection.GetAppConfig()
	connString := env.GetDBEnv()

	isConfigProvided := config.IsConfigFileProvided
	if isConfigProvided {
		connString = config.PostgresDB
	}

	connectWithConnString(connString)
}

func connectWithConnString(conn string) {
	logger := logger.NewLogger()
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		logger.ErrorLogger.Printf("Couldn't connect to Database %+v", err)
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	logger.InfoLogger.Printf("Successfully connected to Database : %+v", conn)

	ORM = db
}
