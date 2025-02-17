package postgres

import (
	"App/internal/config"

	"fmt"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPosgrs(cfg config.DbCOnfig) (*gorm.DB, error) {
	dns := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"

	database, err := gorm.Open(pg.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("NewPostgres: %w", err)
	}

	return database, nil

}
