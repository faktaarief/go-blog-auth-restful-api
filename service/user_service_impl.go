package service

import (
	"github.com/faktaarief/go-blog-auth-restful-api/helper"
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

func (s *userService) Create(userRequest helper.UserRequest) (domain.User, error) {
	user := domain.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}
