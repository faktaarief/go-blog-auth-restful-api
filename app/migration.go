package app

import (
	"github.com/faktaarief/go-blog-auth-restful-api/model/domain"
	"gorm.io/gorm"
)

func MigrationDB() *gorm.DB {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
