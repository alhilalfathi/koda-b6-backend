package main

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/middleware"
	"koda-b6-backend/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// @title Backend Coffeeshop
// @version 1.0.0
// @description this is basic bakcend apps with DI
// @host localhost:8888
// @BasePath /
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

	routes.SetupRoutes(r, conn)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
