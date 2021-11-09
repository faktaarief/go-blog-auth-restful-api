package repository

import "github.com/faktaarief/go-blog-auth-restful-api/model/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (domain.User, error)
}
