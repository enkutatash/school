package middleware

import (
	"fmt"
	"net/http"
	schoolerrors "schoolbackend/errors"
	"schoolbackend/token"
	"strings"

	"github.com/gin-gonic/gin"
)
func AuthenticateParent() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("Parent Middleware")
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken ==""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": schoolerrors.ErrorUnauthorizedAccess.Message})			
			c.Abort()
			return
		}
		splitToken := strings.Split(clientToken, "Bearer ")
		if len(splitToken) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": schoolerrors.ErrorInvalidHeaderFormat.Message})
            c.Abort()
            return
        }
		clientToken = splitToken[1]

        claims, err := token.ValidateToken(clientToken)
        if err != "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": schoolerrors.ErrorInvalidToken.Message})
            c.Abort()
            return
        }
		userRole := claims.Role 
        if userRole != "parent" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": schoolerrors.ErrorParentAccessOnly.Message})
            c.Abort()
            return
        }
		c.Next()
    }
}