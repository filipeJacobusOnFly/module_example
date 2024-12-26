package controllers

import (
	"net/http"
	"strings"

	repositories "module_example/src/http/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(tokenRepo repositories.TokenRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logrus.Warn("Authorization header is required")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenValue := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenValue == authHeader {
			logrus.Warn("Invalid token format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token, err := tokenRepo.GetToken(tokenValue)
		if err != nil || token == nil {
			logrus.Warn("Invalid or expired token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		logrus.Infof("Token validated successfully: %s", tokenValue)
		c.Next()
	}
}
