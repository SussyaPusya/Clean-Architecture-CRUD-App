package postgres

import (
	"App/internal/config"
	"database/sql"
	"fmt"
)

func NewPosgrs(cfg config.DbCOnfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("NewPostgres: %w", err)
	}

	return database, nil

}
