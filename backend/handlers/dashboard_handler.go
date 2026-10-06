package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	dockerService "secure-login-system/docker"
)

func Dashboard(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Role not found",
		})
		return
	}

	roleString, ok := role.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid role",
		})
		return
	}

	switch roleString {
	case "admin":
		if err := dockerService.StartNginx(); err != nil {
			log.Printf("failed to start Nginx: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Nginx failed to start",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Welcome to Admin Dashboard",
			"role":         "admin",
			"nginx_status": "running",
			"nginx_url":    "http://localhost:8081",
		})

	case "user":
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to User Dashboard",
			"role":    "user",
		})

	default:
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Unknown role",
		})
	}
}
