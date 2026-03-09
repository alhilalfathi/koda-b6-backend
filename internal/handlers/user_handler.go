package handlers

import (
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(sr *service.UserService) *UserHandler {
	return &UserHandler{
		service: sr,
	}
}

// get all user
func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input invalid",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List of all users",
		"results": users,
	})
}

// get user by id
func (h *UserHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.service.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input Invalid",
		})
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User found",
		"results": user,
	})
}

// create user
func (h *UserHandler) Create(ctx *gin.Context) {
	var newUser models.CreateUserRequest

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Create user failed",
		})
		return
	}
	h.service.Create(&newUser)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": true,
		"message": "Create user successfuly",
	})
}

// update user
func (h *UserHandler) Update(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email cannot blank",
		})
		return
	}

	var user models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	updatedUser, err := h.service.Update(email, &user)
	if err != nil {
		if err.Error() == "User not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}
		if err.Error() == "Password cannot blank" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated successfully",
		"results": updatedUser,
	})
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	err := h.service.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": true,
		"message": "User delete successfully",
	})
}
