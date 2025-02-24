package repository

import (
	"App/internal/entity"
	errogo "App/pkg/erro.go"
	"sync"
)

type Repository struct {
	Redis map[string]*entity.User
	mu    sync.Mutex
}

func NewRepository() *Repository {

	return &Repository{Redis: make(map[string]*entity.User), mu: sync.Mutex{}}
}

func (r *Repository) GetData(user entity.User) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	data, ok := r.Redis[user.UserName]
	if ok {
		return data, nil
	}

	return &entity.User{}, errogo.ErrUsrNotFOund

}

func (r *Repository) CreateData(user entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.Redis[user.UserName]
	if ok {
		return errogo.ErrUsrREg
	}
	r.Redis[user.UserName] = &user
	return nil

}

func (r *Repository) UpdateData(user entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.Redis[user.UserName]
	if ok {
		r.Redis[user.UserName] = &user

		return nil
	}

	return errogo.ErrUsrNotFOund

}

func (r *Repository) DeleteData(user entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.Redis[user.UserName]

	if ok {
		delete(r.Redis, user.UserName)

		return nil
	}

	return errogo.ErrUsrNotFOund

}
