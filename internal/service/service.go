package service

import (
	"App/internal/entity"
	"fmt"
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

	err := s.repo.CreateData(user)
	if err != nil {
		log.Fatal("Err: never create data", err)
	}

	return nil

}

func (s *Service) GetData(user entity.User) (*entity.User, error) {

	person, err := s.repo.GetData(user)
	if err != nil {
		return &entity.User{}, fmt.Errorf("zalupa %e", err)

	}

	return person, nil

}
func (s *Service) UpdateData(user entity.User) error {

	err := s.repo.UpdateData(user)
	if err != nil {
		return fmt.Errorf("zalupa %e", err)

	}
	return nil

}
func (s *Service) DeleteData(user entity.User) error {

	err := s.repo.DeleteData(user)
	if err != nil {
		return fmt.Errorf("zalupa %e", err)

	}

	return nil

}
