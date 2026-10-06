package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"secure-login-system/utils"
)

type AuthHandler struct {
	DB        *sql.DB
	JWTSecret string
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Enter a valid email and password",
		})
		return
	}

	request.Email = strings.TrimSpace(request.Email)

	var userID int
	var passwordHash string
	var role string

	err := h.DB.QueryRow(
		"SELECT id, password_hash, role FROM users WHERE email = ?",
		request.Email,
	).Scan(&userID, &passwordHash, &role)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPassword(request.Password, passwordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateToken(h.JWTSecret, userID, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
