package handlers

import (
	"koda-b6-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var data models.Users
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Register Failed",
		})
		return
	}
	if data.Email == "" || data.Password == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email and Password cannot blank",
		})
		return
	}

	for i := 0; i < len(models.UserList); i++ {
		if models.UserList[i].Email == data.Email {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Email already exist",
			})
			return
		}
	}

	hashedPassword, err := HashPassword(data.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to hash password",
		})
		return
	}

	data.Password = hashedPassword
	data.Id = models.NextId
	models.UserList = append(models.UserList, data)
	models.NextId++
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Register Success",
	})
}

func Login(ctx *gin.Context) {
	var data models.Users
	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Login Failed",
		})
		return
	}
	if data.Email == "" || data.Password == "" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email and Password cannot blank",
		})
		return
	}

	for i := 0; i < len(models.UserList); i++ {
		if models.UserList[i].Email == data.Email {
			if VerifyPassword(models.UserList[i].Password, data.Password) {
				ctx.JSON(http.StatusOK, models.Response{
					Success: true,
					Message: "Login successful",
				})
			} else {
				ctx.JSON(http.StatusBadRequest, models.Response{
					Success: false,
					Message: "Password Incorrect",
				})
			}
			return
		}
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Email Incorrect",
		})

	}
}
