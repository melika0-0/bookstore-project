package db

import (
	"github.com/melika0-0/bookstore-project/api/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
//this is for connection
func InitDB(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s ,port=%d, user=%s, password=%s, sslmode=%s ,dbname=%s , TimeZone=Asia/Tehran",
             		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.SSLMode, cfg.Postgres.DbName)
    dbClient , err := gorm.Open(postgres.Open(cnn), &gorm.Config{}) //we need a driver 
	if err != nil {
		return err
	}
	sqlDb , _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(cfg.Postgres.Maxidleconnections)
	sqlDb.SetMaxOpenConns(cfg.Postgres.Maxopenconnections)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.Connmaxtimeout * time.Minute)
	log.Println("connected to postgress")
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	conection , _ := dbClient.DB()
	conection.Close()
}

