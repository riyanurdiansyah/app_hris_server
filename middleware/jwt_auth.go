package middleware

import (
	"app-hris-server/helper"
	"app-hris-server/service"
	"net/http"
	"strings"

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
		if strings.Contains(authHeader, "Bearer") {
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 0 {
				token, err := jwtService.ValidateToken(splitToken[1])

				if err != nil {
					response := helper.DefaultErrorResponse{
						Code:    http.StatusUnauthorized,
						Message: "Unauthorized - " + err.Error(),
					}
					c.AbortWithStatusJSON(http.StatusUnauthorized, response)
					return
				}
				if !token.Valid {
					response := helper.DefaultErrorResponse{
						Code:    http.StatusUnauthorized,
						Message: "Unauthorized - " + err.Error(),
					}
					c.AbortWithStatusJSON(http.StatusUnauthorized, response)
					return
				}
			} else {
				response := helper.DefaultErrorResponse{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized - Token is not found",
				}
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		} else {
			response := helper.DefaultErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized - Token must be Bearer",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}
}
