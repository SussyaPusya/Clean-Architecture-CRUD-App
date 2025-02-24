package repository_test

import (
	"App/internal/entity"
	"App/internal/repository"
	"fmt"
	"testing"
)

func TestCreateData(t *testing.T) {
	botCreate := entity.User{
		UserName: "NAZIK",
		Password: "asfdasfsa",
		BirhDay:  "anal",
	}

	repo := repository.NewRepository()

	err := repo.CreateData(botCreate)

	if err != nil {
		t.Error("dasfasf", err)
	}

}

func TestGetData(t *testing.T) {

	bobGet := entity.User{
		UserName: "NAZIK",
	}

	botCreate := entity.User{
		UserName: "NAZIK",
		Password: "asfdasfsa",
		BirhDay:  "anal",
	}

	repo := repository.NewRepository()

	err := repo.CreateData(botCreate)

	if err != nil {
		t.Error("dasfasf", err)
	}

	get, err := repo.GetData(bobGet)
	if err != nil {
		t.Error("dasfasadsaf", err)
	}

	fmt.Println(get, botCreate)

}

func TestUpdateData(t *testing.T) {
	botCreate := entity.User{
		UserName: "NAZIK",
		Password: "asfdasfsa",
		BirhDay:  "anal",
	}

	bobUpdate := entity.User{
		UserName: "NAZIK",
		Password: "12312",
		BirhDay:  "COck",
	}

	repo := repository.NewRepository()

	err := repo.CreateData(botCreate)

	if err != nil {
		t.Error("dasfasf", err)
	}

	err = repo.UpdateData(bobUpdate)
	if err != nil {
		t.Error("CAAAL")
	}

}

func TestDeleteData(t *testing.T) {
	botCreate := entity.User{
		UserName: "NAZIK",
		Password: "asfdasfsa",
		BirhDay:  "anal",
	}
	bobGet := entity.User{
		UserName: "NAZIK",
	}

	repo := repository.NewRepository()

	err := repo.CreateData(botCreate)

	if err != nil {
		t.Error("dasfasf", err)
	}

	err = repo.DeleteData(bobGet)
	if err != nil {
		t.Error("CAAAL")
	}

}
