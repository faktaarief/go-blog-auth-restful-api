package repository

import (
	"github.com/faktaarief/go-blog-auth-restful-api/model/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]domain.User, error) {
	var users []domain.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindById(id int) (domain.User, error) {
	var user domain.User

	err := r.db.Find(&user, id).Error

	return user, err
}

func (r *repository) Create(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Update(userModel, userRequest domain.User) (domain.User, error) {
	err := r.db.Model(&userModel).Updates(userRequest).Error

	return userModel, err
}

func (r *repository) Delete(user domain.User) (domain.User, error) {
	err := r.db.Unscoped().Delete(&user).Error

	return user, err
}
