package service

import (
	"App/internal/entity"
	utils "App/pkg"
	erro "App/pkg/erro.go"
	"log"
)

type Repository interface {
	CreateData(entity.User) error
	GetData(entity.User) (*entity.User, error)
	UpdateData(entity.User) error
	DeleteData(entity.User) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateData(user entity.User) error {
	if utils.ValidAuth(user) {
		return erro.ErrAuth
	}

	err := s.repo.CreateData(user)
	if err != nil {
		log.Fatal("Err: never create data")
	}

	return nil

}

func (s *Service) GetData(user entity.User) (*entity.User, error) {
	if utils.ValidAuth(user) {
		person, err := s.repo.GetData(user)
		if err != nil {

		}

		return person, nil

	}
	return nil, erro.ErrUnauth

}
func (s *Service) UpdateData(user entity.User) error {
	if utils.ValidAuth(user) {
		err := s.repo.UpdateData(user)
		if err != nil {

		}

		return nil

	}
	return erro.ErrUnauth

}
func (s *Service) DeleteData(user entity.User) error {
	if utils.ValidAuth(user) {
		err := s.repo.DeleteData(user)
		if err != nil {

		}

		return nil

	}
	return erro.ErrUnauth

}
