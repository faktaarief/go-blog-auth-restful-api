package service

import (
	"github.com/faktaarief/go-blog-auth-restful-api/helper"
	"github.com/faktaarief/go-blog-auth-restful-api/model/domain"
)

type UserService interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (domain.User, error)
	Create(user helper.UserRequest) (domain.User, error)
	Update(id int, user helper.UserRequest) (domain.User, error)
	Delete(id int) (domain.User, error)
}
