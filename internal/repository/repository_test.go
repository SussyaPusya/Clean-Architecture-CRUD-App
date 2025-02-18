package repository_test

import (
	"App/internal/config"
	"App/internal/entity"
	"App/internal/repository"
	"App/pkg/postgres"
	"fmt"
	"testing"
)

func TestCreateData(t *testing.T) {
	db, err := postgres.NewPosgrs(config.NewConfig().DbCfg)
	if err != nil {
		t.Error("Кал дб не робит")
	}

	repo := repository.NewRepository(db)

	Bob := entity.User{
		UserName: "Nazi1k",
		Password: "1212334",
		Auth:     "553213",
		BirhDay:  "1231223",
	}

	err = repo.CreateData(Bob)
	if err != nil {
		t.Error("кал ошибка ", err)
	}

}

func TestGetData(t *testing.T) {
	db, err := postgres.NewPosgrs(config.NewConfig().DbCfg)
	if err != nil {
		t.Error("Кал дб не робит")
	}

	repo := repository.NewRepository(db)

	bob := entity.User{
		UserName: "Jhane",
		Password: "Doe",
		Auth:     "553213",
		BirhDay:  "11.23.12",
	}

	pers, err := repo.GetData(bob)
	if err != nil {
		t.Error("cal")
	}
	fmt.Println(pers)

}
