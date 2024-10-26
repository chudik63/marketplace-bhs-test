package middleware

import (
	"log"
	"marketplace-bhs-test/internal/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenManager *auth.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty authorization token"})
			c.Abort()
			return
		}

		userId, exp, err := tokenManager.Parse(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token"})
			c.Abort()
			return
		}
		log.Print(userId)
		if float64(time.Now().Unix()) > exp {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		c.Set("userID", uint64(userId))
		c.Next()
	}
}
