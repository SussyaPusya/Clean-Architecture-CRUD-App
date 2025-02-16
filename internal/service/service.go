package service

import (
	"App/internal/entity"
	utils "App/pkg"
	erro "App/pkg/erro.go"
)

type Repository interface {
	CreateData(entity.User) (*entity.User, error)
	GetData(entity.User) (*entity.User, error)
	UpdateData(entity.User) (*entity.User, error)
	DeleteData(entity.User) (*entity.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateData(user entity.User) (*entity.User, error) {
	if utils.ValidAuth(user) {
		return nil, erro.ErrAuth
	}

	person, err := s.repo.CreateData(user)
	if err != nil {

	}

	return person, nil

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
func (s *Service) UpdateData(user entity.User) (*entity.User, error) {
	if utils.ValidAuth(user) {
		person, err := s.repo.UpdateData(user)
		if err != nil {

		}

		return person, nil

	}
	return nil, erro.ErrUnauth

}
func (s *Service) DeleteData(user entity.User) (*entity.User, error) {
	if utils.ValidAuth(user) {
		person, err := s.repo.DeleteData(user)
		if err != nil {

		}

		return person, nil

	}
	return nil, erro.ErrUnauth

}
