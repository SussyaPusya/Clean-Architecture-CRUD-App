package repository_test

import (
	"App/internal/config"
	"App/internal/entity"
	"App/internal/repository"
	"App/pkg/postgres"
	"testing"
)

func TestCreateData(t *testing.T) {
	db, err := postgres.NewPosgrs(config.NewConfig().DbCfg)
	if err != nil {
		t.Error("Кал дб не робит")
	}

	repo := repository.NewRepository(db)

	Bob := entity.User{
		UserName: "fashik",
		Password: "sdsadsaa",
		Auth:     "yeasdsdss",
		BirhDay:  "ds12312a",
	}

	_, err = repo.CreateData(Bob)
	if err != nil {
		t.Error("кал ошибка ", err)
	}

}
