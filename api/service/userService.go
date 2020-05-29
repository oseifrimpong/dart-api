package service

import (
	"dart-api/api/model"
	"dart-api/api/repository"
)

type Service interface {
	Create(u *model.User) error
	Update(u *model.User) error
	GetByID(id string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Delete(id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) Service {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) Create(u *model.User) error { return svc.repo.Create(u) }

func (svc *userService) GetAll() ([]*model.User, error) { return svc.repo.GetAll() }

func (svc *userService) GetByID(id string) (*model.User, error) { return svc.repo.GetByID(id) }

func (svc *userService) Delete(id string) error { return svc.repo.Delete(id) }

func (svc *userService) Update(u *model.User) error { return svc.repo.Update(u) }
