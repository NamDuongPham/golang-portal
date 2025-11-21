package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/internal/validators"
)

type AdminHandler struct {
	AuthService services.AuthServiceInterface
}

// LOGIN
func (h *AdminHandler) Login(c *gin.Context) {
	body, ok := validators.ValidateLogin(c)
	if !ok {
		return
	}

	res := h.AuthService.Login(body.Username, body.Password)
	if !res.Status {
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	c.JSON(http.StatusOK, res)
}

// LOGOUT
func (h *AdminHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]

	res := h.AuthService.Logout(token)
	if !res.Status {
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	c.JSON(http.StatusOK, res)
}

// REFRESH
func (h *AdminHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	res := h.AuthService.RefreshToken(req.RefreshToken)
	if !res.Status {
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	c.JSON(http.StatusOK, res)
}
