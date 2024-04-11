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

	if config.IsConfigFileProvided {
		connString = injection.GetDBString()
	}

	connectWithConnString(connString)
}

func connectWithConnString(conn string) {
	logger := logger.NewLogger()
	db, err := gorm.Open("postgres", conn)
	// db, err := gorm.Open("postgres", "tiger:tigercoders@tiger-db")
	if err != nil {
		logger.ErrorLogger.Printf("Couldn't connect to Database| error: %+v | db conn string: %+v", err, conn)
		log.Fatalf("Error connectiong to Database | error: %+v | db conn string: %+v", err, conn)
	}
	logger.InfoLogger.Printf("Successfully connected to Database : %+v", conn)

	ORM = db
}
