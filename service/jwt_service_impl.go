package service

import (
	"app-ecommerce-server/helper"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		secretKey: "riyanurdiansyah",
		issuer:    getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "riyan"
	}
	return secretKey
}

func (service *JWTServiceImpl) GenerateToken(UserId string, Email string) string {
	claims := &JWTCustomClaim{
		UserId: UserId,
		Email:  Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Now().Hour() + 1)).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	helper.PanicIfError(err)
	return t
}

func (service *JWTServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
