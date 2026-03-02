package main

import (
	"context"
	"fmt"
	"koda-b6-backend/handlers"
	"koda-b6-backend/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

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

	godotenv.Load()

	connConfig, err := pgx.ParseConfig("")

	if err != nil {
		fmt.Println("Failed to parse config")
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	if err != nil {
		fmt.Println("Failed to connecting db")
	}

	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/register", handlers.Register)

	r.POST("/login", handlers.Login)

	r.GET("/users", func(ctx *gin.Context) {

		rows, err := conn.Query(context.Background(), `
		SELECT id, email, password
		FROM "USER"
		`)

		users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Failed to get data users",
			})
			return
		}

		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "List of users",
			Results: users,
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

	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
