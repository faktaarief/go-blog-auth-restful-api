package handler

import (
	"net/http"
	"strconv"

	"github.com/faktaarief/go-blog-auth-restful-api/helper"
	"github.com/faktaarief/go-blog-auth-restful-api/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) FindAll(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	userResponses := helper.UserResponses(users)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponses,
	})
}

func (h *userHandler) FindById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.FindById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	userResponse := helper.UserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}
