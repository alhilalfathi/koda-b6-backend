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
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input invalid",
		})
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List of all users",
		Results: users,
	})
}

// get user by id
func (h *UserHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.service.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input Invalid",
		})
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User found",
		Results: user,
	})
}

// get user by email
func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := h.service.GetById(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input Invalid",
		})
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User found",
		Results: user,
	})
}

// create user
func (h *UserHandler) Create(ctx *gin.Context) {
	var newUser models.CreateUserRequest

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Create user failed",
		})
		return
	}

	if err := h.service.Register(&newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Register Failed",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, models.Response{
		Success: true,
		Message: "Register successfuly",
	})
}

// user login
func (h *UserHandler) Login(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email and password required",
		})
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login success",
		Results: models.LoginUserResponse{Token: token},
	})
}

// update user
func (h *UserHandler) Update(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email cannot blank",
		})
		return
	}

	var user models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	updatedUser, err := h.service.Update(email, &user)
	if err != nil {
		if err.Error() == "User not found" {
			ctx.JSON(http.StatusNotFound, models.Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		if err.Error() == "Password cannot blank" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User updated successfully",
		Results: updatedUser,
	})
}

// delete user
func (h *UserHandler) Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	err := h.service.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, models.Response{
		Success: true,
		Message: "User delete successfully",
	})
}
