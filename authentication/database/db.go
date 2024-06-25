package database

import (
	"database/sql"
	"pesto-auth/log"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func Init() {
	connStr := "postgresql://pesto:pest02024@localhost:5432/pesto?sslmode=disable"
	// Connect to database
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Logger.Error("failed operning connection to DB", zap.Any("error", err))
	}
	log.Logger.Debug("DB connection initialised")
	log.Logger.Info("Running DB migrations...")
	RunMigrations()
}

func HandleDBclose() int {
	DB.Close()
	return 1
}

func RunMigrations() {
	_, err := DB.Query(`CREATE TABLE IF NOT EXISTS users 
	(id UUID PRIMARY KEY, name VARCHAR(200) NULL,	email VARCHAR(200) NULL,
	password VARCHAR(200) NULL, country VARCHAR(64) NULL, phone VARCHAR(20) NULL)`)
	if err != nil {
		log.Logger.Error("Failed running migrations : ", zap.Any("error", err))
	}
}
