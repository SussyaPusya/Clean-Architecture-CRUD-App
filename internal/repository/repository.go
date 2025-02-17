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

// func (r *Repository) GetData(user entity.User) (*entity.User, error) {

// }

func (r *Repository) CreateData(user entity.User) (*entity.User, error) {

	usa := utils.TransformStruct(user)

	result := r.DB.Create(&usa)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	var bob entity.UserDb
	r.DB.First(&bob)
	fmt.Println(&bob)
	person := utils.ReversTransStruct(bob)
	return person, nil

}

// func (r *Repository) UpdateData(user entity.User) (*entity.User, error) {

// }
// func (r *Repository) DeleteData(user entity.User) (*entity.User, error) {

// }
