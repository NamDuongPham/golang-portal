package http

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
)

type AdminHandler struct {
	AuthService services.AuthServiceInterface
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AdminHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid data"})
		return
	}

	token, refreshToken, err := h.AuthService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"message":       "Login success",
			"access_token":  token,
			"refresh_token": refreshToken,
		},
	})
}
func (h *AdminHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "missing access token"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	err := h.AuthService.Logout(tokenString)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Logout successful"})
}
