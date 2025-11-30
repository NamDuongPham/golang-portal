package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/services"
	"github.com/namduong/project-layout/internal/validators"
)

type PortalHandler struct {
	AuthPortalService services.AuthPortalServiceInterface
}

// LOGIN
func (h *PortalHandler) Login(c *gin.Context) {
	body, ok := validators.ValidateLogin(c)
	if !ok {
		return
	}
	res := h.AuthPortalService.Login(body.Username, body.Password)
	if !res.Status {
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	c.JSON(http.StatusOK, res)
}

// LOGOUT
func (h *PortalHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := authHeader[len("Bearer "):]

	res := h.AuthPortalService.Logout(token)
	if !res.Status {
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	c.JSON(http.StatusOK, res)
}

// REFRESH
func (h *PortalHandler) RefreshToken(c *gin.Context) {
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

	res := h.AuthPortalService.RefreshToken(req.RefreshToken)
	if !res.Status {
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	c.JSON(http.StatusOK, res)
}
