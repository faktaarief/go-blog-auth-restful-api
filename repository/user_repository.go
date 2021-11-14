package repository

import "github.com/faktaarief/go-blog-auth-restful-api/model/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(userModel, userRequest domain.User) (domain.User, error)
	Delete(user domain.User) (domain.User, error)
}
