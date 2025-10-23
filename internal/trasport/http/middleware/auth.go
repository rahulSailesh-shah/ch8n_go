package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/rahulSailesh-shah/ch8n_go/pkg/auth"
)

func AuthMiddleware(authKeys jwk.Set) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := auth.UserFromJWT(c.Request, authKeys)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
