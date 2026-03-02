package main

import (
	"fmt"
	"koda-b6-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

func HashPassword(password string) (string, error) {
	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(encodedHash, password string) bool {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(encodedHash))
	if err != nil {
		return false
	}
	return ok
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:5432")
		ctx.Header("Access-Control-Allow-Headers", "Content-type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}

	}
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/register", func(ctx *gin.Context) {
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
	})

	r.POST("/login", func(ctx *gin.Context) {
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
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "List of users",
			Results: models.UserList,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i := range models.UserList {
			if fmt.Sprint(models.UserList[i].Id) == id {
				ctx.JSON(http.StatusOK, models.Response{
					Success: true,
					Message: fmt.Sprintf("Hello User %s", id),
					Results: models.UserList[i],
				})
				return
			}
		}
		ctx.JSON(404, models.Response{
			Success: false,
			Message: fmt.Sprintf("User %s not found", id),
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var newData models.Users
		err := ctx.ShouldBind(&newData)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Input error",
			})
			return
		}
		for i := range models.UserList {
			if fmt.Sprint(models.UserList[i].Id) == id {
				if newData.Email != "" {
					for j := range models.UserList {
						if models.UserList[j].Email == newData.Email && models.UserList[j].Id != newData.Id {
							ctx.JSON(http.StatusBadRequest, models.Response{
								Success: false,
								Message: "Email already registered",
							})
							return
						}
					}
					models.UserList[i].Email = newData.Email
				}
				if newData.Password != "" {
					models.UserList[i].Password = newData.Password
				}
				ctx.JSON(http.StatusOK, models.Response{
					Success: true,
					Message: "User data updated",
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Input error",
			})
		}
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "User not found",
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i := range models.UserList {
			if fmt.Sprint(models.UserList[i].Id) == id {
				models.UserList = append(models.UserList[:i], models.UserList[i+1:]...)
				ctx.JSON(http.StatusOK, models.Response{
					Success: true,
					Message: "User deleted",
				})
			}
		}
	})

	r.Run("localhost:8888")
}
