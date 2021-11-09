package main

import (
	"fmt"

	"github.com/faktaarief/go-blog-auth-restful-api/app"
	"github.com/faktaarief/go-blog-auth-restful-api/handler"
	"github.com/faktaarief/go-blog-auth-restful-api/repository"
	"github.com/faktaarief/go-blog-auth-restful-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := app.MigrationDB()

	repository.NewUserRepository(db)

	fmt.Println("Database Successfully Connected")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users", userHandler.FindAll)

	router.Run(":3000")
}
