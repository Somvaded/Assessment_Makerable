package middlewares

import (
	"net/http"

	"github.com/Somvaded/assessment/utils"
	"github.com/gin-gonic/gin"
)

func Protect() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString ,err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing auth_token cookie"})
			c.Abort()
			return
		}
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func CheckRole(role string) gin.HandlerFunc{
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		userRole := roleVal.(string)
		if userRole == role {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}

