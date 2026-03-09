package main

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/di"
	"koda-b6-backend/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

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

	r.Use(middleware.CorsMiddleware())

	container := di.NewContainer(conn)

	userHandler := container.UserHandler()
	productHandler := container.ProductHandler()

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/register", userHandler.Create)
		auth.POST("/login", userHandler.Login)
	}

	users := r.Group("/users")
	{
		users.GET("/", userHandler.GetAll)
		users.GET("/:id", userHandler.GetById)
		users.GET("/:email", userHandler.GetByEmail)
		users.PATCH("/:email", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	products := r.Group("/product")
	{
		products.POST("/", productHandler.CreateProduct)
		products.GET("/", productHandler.GetAllProduct)
		products.GET("/:id", productHandler.GetProductById)
		products.PATCH("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)

	}

	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
