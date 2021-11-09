package service

import "github.com/faktaarief/go-blog-auth-restful-api/model/domain"

type UserService interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (domain.User, error)
}
