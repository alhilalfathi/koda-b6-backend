package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		godotenv.Load()
		ctx.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-type, authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}
	}
}
