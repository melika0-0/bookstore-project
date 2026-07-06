package db

import (
	"fmt"
	"log"
	"time"

	"github.com/melika0-0/bookstore-project/api/config"
	"github.com/melika0-0/bookstore-project/data/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbClient *gorm.DB

func InitDB(cfg *config.Config) error {
//dsn is urrrr connection string 
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)

	// Connect to PostgreSQL
	db, err := gorm.Open(
	postgres.Open(dsn),
	&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	},)
	if err != nil {
		return err
	}

	// Run migrations
	if err := db.AutoMigrate(
		&models.Book{},
	); err != nil {
		return err
	}

	// Get underlying sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(
		time.Duration(cfg.Postgres.ConnMaxTimeout) * time.Minute,
	)

	dbClient = db

	log.Println("connected to postgres")
	log.Println("book migration completed")

	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	if dbClient == nil {
		return
	}

	sqlDB, err := dbClient.DB()
	if err != nil {
		return
	}

	_ = sqlDB.Close()

	log.Printf("SSL MODE = '%s'", config.GetConfig().Postgres.SSLMode)
}