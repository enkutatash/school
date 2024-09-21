package middleware

import (
	"net/http"
	"schoolbackend/token"
	"strings"

	"github.com/gin-gonic/gin"
)
func AuthenticateParent() gin.HandlerFunc {
    return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken ==""{
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is required"})
			c.Abort()
			return
		}
		splitToken := strings.Split(clientToken, "Bearer ")
		if len(splitToken) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }
		clientToken = splitToken[1]

        claims, err := token.ValidateToken(clientToken)
        if err != "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
		userRole := claims.Role 
        if userRole != "student" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Only Students have this access"})
            c.Abort()
            return
        }
		c.Next()
    }
}