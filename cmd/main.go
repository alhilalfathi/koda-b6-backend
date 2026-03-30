package main

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/middleware"
	"koda-b6-backend/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// @title Backend Coffeeshop
// @version 1.0.0
// @description this is basic bakcend apps with DI
// @host localhost:8888
// @BasePath /
func main() {

	godotenv.Load()

	// connConfig, err := pgx.ParseConfig("")
	// if err != nil {
	// 	fmt.Println("Failed to parse config")
	// }

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
	)

	// config, err := pgxpool.ParseConfig(dbURL)
	// if err != nil {
	// 	fmt.Printf("Failed to parse config: %v\n", err)
	// 	os.Exit(1)
	// }

	// config.MaxConns = 20
	// config.MinConns = 5

	//conn, err := pgx.Connect(context.Background(), connConfig.ConnString()) //bisa pke pgxpool biar ga conn busy
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Println("Failed to connecting db")
		os.Exit(1)
	}

	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Printf("Failed to ping database: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.Use(middleware.CorsMiddleware())

	// routes.SetupRoutes(r, conn)
	routes.SetupRoutes(r, pool)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
