package database

import (
	"database/sql"
	"pesto-product-manager/log"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB

func Init(cfg Config) {
	connStr := "postgresql://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + "/" + cfg.DBname + "?sslmode=disable"
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
	_, err := DB.Query(`CREATE TABLE IF NOT EXISTS products
	 (id UUID PRIMARY KEY, name VARCHAR(200) NULL,	category VARCHAR(100) NULL,
	 manufacturer VARCHAR(200) NULL, description VARCHAR(200) NULL, price FLOAT NULL, 
	 origin VARCHAR(64) NULL, last_updated DATETIME NULL, created_at DATETIME NULL)`)
	if err != nil {
		log.Logger.Error("Failed running migrations : ", zap.Any("error", err))
		log.Logger.Info("Trying to run migrations again in 2 secs...")
		time.Sleep(2 * time.Second)
		RunMigrations()
	}
	log.Logger.Info("Migrations ran successfully")
}
