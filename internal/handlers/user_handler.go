package handlers

import (
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/service"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(sr *service.UserService) *UserHandler {
	return &UserHandler{
		service: sr,
	}
}

// GetUsers godoc
// @Summary Get all user account
// @Description Show all users
// @Tags User
// @Produce json
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/users [get]
func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to get users",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List of all users",
		Results: users,
	})
}

// GetUsersByID godoc
// @Summary Get user by ID
// @Description Get user detail by ID
// @Tags User
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /admin/users/{id} [get]
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

// GetUsersByEmail godoc
// @Summary Get user by Email
// @Description Get user detail by email
// @Tags User
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/users/email/{email} [get]
func (h *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := h.service.GetByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Internal server error",
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

// UserRegister godoc
// @Summary Register user
// @Description Create new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Register Request"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /auth/register [post]
func (h *UserHandler) Create(ctx *gin.Context) {
	var newUser models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
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
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Register successfuly",
	})
}

// UserLogin godoc
// @Summary Login user
// @Description Login user and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginUserRequest true "Login Request"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /auth/login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email and password required",
		})
		return
	}

	user, token, err := h.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login success",
		Results: models.LoginUserResponse{
			Token:    token,
			Fullname: user.FullName,
		},
	})
}

// UpdateUser (Forgot Password) godoc
// @Summary Update user by email
// @Description Update user data (forgot password flow)
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "User Email"
// @Param request body models.UpdateUserRequest true "Update User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
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

// UpdateProfile godoc
// @Summary Update profile
// @Description Update logged in user profile
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.UpdateUserRequest true "Update Profile"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/users/profile [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userIdRaw, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
			Results: nil,
		})
		return
	}

	userId, ok := userIdRaw.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "invalid userId type",
			Results: nil,
		})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
			Results: nil,
		})
		return
	}

	user := models.Users{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	updatedUser, err := h.service.UpdateById(userId, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
			Results: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User delete successfully",
		Results: updatedUser,
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by email
// @Tags User
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /admin/users/{email} [delete]
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
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User delete successfully",
	})
}

// ChangePassword godoc
// @Summary Change password
// @Description Change password for logged in user
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.ChangePasswordRequest true "Change Password"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /admin/profile/password [put]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userIdRaw, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
			Results: nil,
		})
		return
	}

	userId := int(userIdRaw.(int))

	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
			Results: nil,
		})
		return
	}

	err := h.service.ChangePassword(userId, req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
			Results: nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Password update successfully",
	})
}

// GetProfile godoc
// @Summary Get profile
// @Description Get logged in user profile
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /admin/user/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userIdRaw, _ := c.Get("userId")
	userId := int(userIdRaw.(int))

	user, err := h.service.GetProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Get profile successfully",
		Results: user,
	})
}

// UploadProfilePhoto godoc
// @Summary Upload profile photo
// @Description Upload user profile picture
// @Tags Users
// @Accept multipart/form-data
// @Produce json
// @Param picture formData file true "Profile picture"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /admin/users/profile/photo [patch]
// @Security BearerAuth
func (h *UserHandler) UploadProfilePhoto(ctx *gin.Context) {

	// ambil file
	file, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "File is required",
		})
		return
	}

	// sanitize filename
	safeName := filepath.Base(file.Filename)

	// generate nama unik
	filename := fmt.Sprintf("uploads/%d_%s_%s",
		time.Now().UnixNano(),
		uuid.New().String(),
		safeName,
	)

	// simpan file
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to save file",
		})
		return
	}

	// ambil user id dari middleware
	userId := ctx.MustGet("userId").(int)

	// update ke database
	err = h.service.UpdateProfilePicture(userId, filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Upload success",
		Results: filename,
	})
}
