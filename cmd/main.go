package main

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/cache"
	"koda-b6-backend/internal/lib"
	"koda-b6-backend/internal/middleware"
	"koda-b6-backend/internal/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	// LOAD ENV
	godotenv.Load()

	// DATABASE SETUP
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
	)

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		fmt.Printf("Failed to parse config: %v\n", err)
		os.Exit(1)
	}

	config.MaxConns = 20
	config.MinConns = 5

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Failed to connect database")
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		fmt.Printf("Failed to ping database: %v\n", err)
		os.Exit(1)
	}

	// REDIS SETUP
	rdb := lib.NewRedisClient()

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		fmt.Println("Failed to connect Redis:", err)
		os.Exit(1)
	}

	cache := cache.NewCache(rdb)

	// GIN SETUP
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	// ROUTES + DI
	routes.SetupRoutes(r, pool, cache)

	// RUN SERVER
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	r.Run(":" + port)
}
