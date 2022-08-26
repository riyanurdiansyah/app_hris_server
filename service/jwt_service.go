package service

import "github.com/dgrijalva/jwt-go"

type JWTService interface {
	GenerateToken(userId string, email string) string
	ValidateToken(token string) (*jwt.Token, error)
}
