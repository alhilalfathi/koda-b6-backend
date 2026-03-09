package main

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/di"
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

	container := di.NewContainer(conn)

	userHandler := container.UserHandler()

	r.POST("/register", userHandler.Create)
	r.POST("/login", userHandler.Login)

	users := r.Group("/users")
	{
		users.GET("/", userHandler.GetAll)
		users.GET("/:id", userHandler.GetById)
		users.GET("/:email", userHandler.GetByEmail)
		users.PATCH("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
