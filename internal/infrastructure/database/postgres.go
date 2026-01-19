package database

import (
	"fmt"

	"github.com/vitaly06/shop-rest-api/internal/infrastructure/config"
	"github.com/vitaly06/shop-rest-api/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db := database

	return db, nil

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&entities.User{},
	)
}
