package middleware

import (
	"marketplace-bhs-test/internal/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenManager *auth.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/users/sign-in" || c.Request.URL.Path == "/users/sign-up" || c.Request.URL.Path == "/users/sign-out" {
			c.Next()
			return
		}

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

		if float64(time.Now().Unix()) > exp {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}

		c.Set("userID", userId)
		c.Next()
	}
}
