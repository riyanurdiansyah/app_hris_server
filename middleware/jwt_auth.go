package middleware

import (
	"app-ecommerce-server/helper"
	"app-ecommerce-server/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.DefaultErrorResponse{
				Code:    401,
				Message: "Unauthorized - Token is not found",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		splitToken := strings.Split(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(splitToken[1])
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			println("USER ID ", claims["user_id"])
			println("ISSUER ", claims["issuer"])
		} else {
			println(err)
			response := helper.DefaultErrorResponse{
				Code:    401,
				Message: "Unauthorized - Token is not valid",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
