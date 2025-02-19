package repository

import (
	"App/internal/entity"
	utils "App/pkg/postgres"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetData(user entity.User) (*entity.User, error) {
	var person entity.UserDb
	usa := utils.TransformStruct(user)

	r.DB.Where(&usa).First(&person)

	user = *utils.ReversTransStruct(person)
	return &user, nil

}

func (r *Repository) CreateData(user entity.User) error {

	usa := utils.TransformStruct(user)

	result := r.DB.Create(&usa)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return nil

}

func (r *Repository) UpdateData(user entity.User) error {
	var person entity.UserDb
	usa := utils.TransformStruct(user)

	r.DB.Where(&usa).First(&person)

	r.DB.Delete(entity.UserDb{}, person.ID)

	pers := utils.ReversTransStruct(*usa)

	r.CreateData(*pers)

	return nil

}

func (r *Repository) DeleteData(user entity.User) error {

	var person entity.UserDb
	usa := utils.TransformStruct(user)

	r.DB.Where(&entity.UserDb{}).First(&person)
	fmt.Println(usa)
	r.DB.Delete(&person)

	return nil

}
