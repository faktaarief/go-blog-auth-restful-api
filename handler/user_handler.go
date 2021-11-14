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

func (h *userHandler) Create(c *gin.Context) {
	var userRequest helper.UserRequest

	/**
	* c.Bind for binding data using content-type "form-data" or "x-www-form-urlencoded" or "JSON"
	* If using "form-" for content-type, we should added tag "form: " in userRequest
	* c.ShouldBindJSON only binding for content-type of JSON
	**/

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	user, err := h.userService.Create(userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *userHandler) Update(c *gin.Context) {
	var userRequest helper.UserRequest

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	/**
	* c.Bind for binding data using content-type "form-data" or "x-www-form-urlencoded" or "JSON"
	* If using "form-" for content-type, we should added tag "form: " in userRequest
	**/

	err := c.Bind(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	user, err := h.userService.Update(id, userRequest)

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

func (h *userHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	user, err := h.userService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	userResponse := helper.UserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete user",
		"data":    userResponse,
	})
}
