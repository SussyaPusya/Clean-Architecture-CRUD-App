package repository

import (
	"App/internal/entity"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetData(user entity.User) (*entity.User, error) {

}

func (r *Repository) CreateData(user entity.User) (*entity.User, error) {

}
func (r *Repository) UpdateData(user entity.User) (*entity.User, error) {

}
func (r *Repository) DeleteData(user entity.User) (*entity.User, error) {

}
