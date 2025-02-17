package postgres

import (
	"App/internal/config"
	"App/internal/entity"

	"fmt"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPosgrs(cfg config.DbCOnfig) (*gorm.DB, error) {

	database, err := gorm.Open(pg.Open(cfg.Dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("NewPostgres: %w", err)
	}

	return database, nil

}

func TransformStruct(user entity.User) *entity.UserDb {
	return &entity.UserDb{UserName: user.UserName, Password: user.Password, BirhDay: user.BirhDay}
}

func ReversTransStruct(user entity.UserDb) *entity.User {
	return &entity.User{UserName: user.UserName, Password: user.Password, BirhDay: user.BirhDay}
}
