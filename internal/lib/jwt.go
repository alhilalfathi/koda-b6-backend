package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type CustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) (string, error) {
	godotenv.Load()
	secret := os.Getenv("APP_SECRET")

	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*CustomClaims, bool) {
	godotenv.Load()
	secret := os.Getenv("APP_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return &CustomClaims{}, false
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return &CustomClaims{}, false
	}

	return claims, true
}
