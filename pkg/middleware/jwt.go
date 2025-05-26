// pkg/middleware/jwt.go
package middleware

import (
	"github.com/teakingwang/gin-demo/pkg/consts"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/pkg/auth"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 把用户信息存入上下文，后续可用
		c.Set(consts.JWTKeyUserID, claims.UserID)
		c.Next()
	}
}
