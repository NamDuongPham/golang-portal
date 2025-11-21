package http

import (
	"net/http"
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

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// ======================= LOGIN =======================
func (h *AdminHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	response := h.AuthService.Login(req.Username, req.Password)
	if !response.Status {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": response.Message,
			"error":   response.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response.Message,
		"data":    response.Data,
	})
}

// ======================= LOGOUT =======================
func (h *AdminHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing access token"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	response := h.AuthService.Logout(tokenString)
	if !response.Status {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": response.Message,
			"error":   response.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// ======================= REFRESH TOKEN =======================
func (h *AdminHandler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if req.RefreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing refresh token"})
		return
	}

	response := h.AuthService.RefreshToken(req.RefreshToken)
	if !response.Status {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": response.Message,
			"error":   response.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response.Message,
		"data":    response.Data,
	})
}
