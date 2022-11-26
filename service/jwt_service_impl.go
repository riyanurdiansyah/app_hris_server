package service

import (
	"app-hris-server/helper"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTCustomClaim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type JWTServiceImpl struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &JWTServiceImpl{
		secretKey: "riyansecret",
		issuer:    getSecretKey(),
	}
}

func getSecretKey() string {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey != "" {
		secretKey = "riyansecret"
	}
	return secretKey
}

func (service *JWTServiceImpl) GenerateToken(UserId string, Email string) string {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	claims := &JWTCustomClaim{
		UserId: UserId,
		Email:  Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			Issuer:    "rsyahproject.co.id",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secretKey))
	helper.PanicIfError(err)
	return t
}

func (service *JWTServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
