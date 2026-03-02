package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

var NextId = 1
var UserList []Users

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
		var data Users
		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Register Failed",
			})
			return
		}
		if data.Email == "" || data.Password == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email and Password cannot blank",
			})
			return
		}

		for i := 0; i < len(UserList); i++ {
			if UserList[i].Email == data.Email {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Message: "Email already exist",
				})
				return
			}
		}

		hashedPassword, err := HashPassword(data.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Failed to hash password",
			})
			return
		}

		data.Password = hashedPassword
		data.Id = NextId
		UserList = append(UserList, data)
		NextId++
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Register Success",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var data Users
		err := ctx.ShouldBind(&data)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Login Failed",
			})
			return
		}
		if data.Email == "" || data.Password == "" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email and Password cannot blank",
			})
			return
		}

		for i := 0; i < len(UserList); i++ {
			if UserList[i].Email == data.Email {
				if VerifyPassword(UserList[i].Password, data.Password) {
					ctx.JSON(http.StatusOK, Response{
						Success: true,
						Message: "Login successful",
					})
				} else {
					ctx.JSON(http.StatusBadRequest, Response{
						Success: false,
						Message: "Password Incorrect",
					})
				}
				return
			}
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Email Incorrect",
			})

		}
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "List of users",
			Results: UserList,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i := range UserList {
			if fmt.Sprint(UserList[i].Id) == id {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: fmt.Sprintf("Hello User %s", id),
					Results: UserList[i],
				})
				return
			}
		}
		ctx.JSON(404, Response{
			Success: false,
			Message: fmt.Sprintf("User %s not found", id),
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var newData Users
		err := ctx.ShouldBind(&newData)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Input error",
			})
			return
		}
		for i := range UserList {
			if fmt.Sprint(UserList[i].Id) == id {
				if newData.Email != "" {
					for j := range UserList {
						if UserList[j].Email == newData.Email && UserList[j].Id != newData.Id {
							ctx.JSON(http.StatusBadRequest, Response{
								Success: false,
								Message: "Email already registered",
							})
							return
						}
					}
					UserList[i].Email = newData.Email
				}
				if newData.Password != "" {
					UserList[i].Password = newData.Password
				}
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "User data updated",
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Input error",
			})
		}
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "User not found",
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i := range UserList {
			if fmt.Sprint(UserList[i].Id) == id {
				UserList = append(UserList[:i], UserList[i+1:]...)
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "User deleted",
				})
			}
		}
	})

	r.Run("localhost:8888")
}
