package database

import (
	"fmt"
	"github.com/yusufbulac/byfood-case/backend/internal/model"
	"log"

	"github.com/yusufbulac/byfood-case/backend/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMySQL(logger *zap.Logger) {
	cfg := config.AppConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	if err := DB.AutoMigrate(&model.Book{}); err != nil {
		logger.Fatal("Database migration failed", zap.Error(err))
	}

	log.Println("Connected to MySQL")
}
