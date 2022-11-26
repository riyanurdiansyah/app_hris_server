package middleware

import (
	"app-hris-server/helper"
	"app-hris-server/service"
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
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized - Token is not found",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		splitToken := strings.Split(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(splitToken[1])
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			println("CHECK ", claims["user_id"])
		} else {
			println(err)
			response := helper.DefaultErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized - " + err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
