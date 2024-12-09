// backend/middlewares/auth.go
package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/utils"
)

func AuthMiddleware(jwtUtil *utils.JWTUtil) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
            c.Abort()
            return
        }

        tokenString := parts[1]
        claims, err := jwtUtil.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // Set user information to context
        c.Set("userID", claims.UserID)
        c.Set("userEmail", claims.Email)

        c.Next()
    }
}
