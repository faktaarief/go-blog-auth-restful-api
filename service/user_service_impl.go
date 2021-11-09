package service

import (
	"github.com/faktaarief/go-blog-auth-restful-api/model/domain"
	"github.com/faktaarief/go-blog-auth-restful-api/repository"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) FindAll() ([]domain.User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *userService) FindById(id int) (domain.User, error) {
	user, err := s.repository.FindById(id)
	return user, err
}
