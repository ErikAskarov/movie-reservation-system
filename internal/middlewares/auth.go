package middlewares

import (
	"movie-reservation-system/internal/helpers"

	"github.com/gin-gonic/gin"
)

func isAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := helpers.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}
