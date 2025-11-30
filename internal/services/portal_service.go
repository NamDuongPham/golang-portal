package services

import (
	"github.com/namduong/project-layout/helper"
	"github.com/namduong/project-layout/internal/auth"
	"github.com/namduong/project-layout/internal/logger"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthPortalService struct {
	PortalRepo       *repositories.PortalRepository
	RefreshTokenRepo *repositories.RefreshTokenRepository
}

func (s *AuthPortalService) saveRefreshToken(rawToken string) error {
	claims, err := auth.DecodeRefreshToken(rawToken)
	if err != nil {
		return err
	}

	rt := &models.RefreshToken{
		ID:        claims.Id,
		Token:     rawToken,
		UserID:    claims.UserID,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	return s.RefreshTokenRepo.Create(rt)
}

func (s *AuthPortalService) Login(username string, password string) helper.Response {
	logger.GetLogger().Info("User login attempt", zap.String("username", username))
	user, err := s.PortalRepo.FindByUsername(username)
	if err != nil {
		return helper.BuildErrorResponse("User not found", "invalid credentials", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return helper.BuildErrorResponse("Wrong password", "invalid credentials", nil)
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, user.UserName)
	if err != nil {
		return helper.BuildErrorResponse("Failed to generate access token", err.Error(), nil)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID, user.UserName)
	if err != nil {
		return helper.BuildErrorResponse("Failed to generate refresh token", err.Error(), nil)
	}

	if err := s.saveRefreshToken(refreshToken); err != nil {
		return helper.BuildErrorResponse("Failed to save refresh token", err.Error(), nil)
	}

	responseData := AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return helper.BuildResponse(true, "Login successful", responseData)
}

func (s *AuthPortalService) Logout(token string) helper.Response {
	claims, err := auth.DecodeAccessToken(token)
	if err != nil {
		return helper.BuildErrorResponse("Invalid access token", err.Error(), nil)
	}
	rt, err := s.RefreshTokenRepo.FindByUserID(claims.UserID)
	if err != nil || rt == nil {
		return helper.BuildErrorResponse("Already logged out", "no active refresh token", nil)
	}
	if err := s.RefreshTokenRepo.DeleteByUserID(claims.UserID); err != nil {
		return helper.BuildErrorResponse("Failed to logout", err.Error(), nil)
	}

	return helper.BuildResponse(true, "Logout successful", nil)
}

func (s *AuthPortalService) RefreshToken(token string) helper.Response {
	claims, err := auth.DecodeRefreshToken(token)
	if err != nil {
		return helper.BuildErrorResponse("Invalid refresh token", err.Error(), nil)
	}

	storedToken, err := s.RefreshTokenRepo.FindByID(claims.Id)
	if err != nil {
		return helper.BuildErrorResponse("Invalid refresh token", "token not found in database", nil)
	}

	if storedToken.Token != token {
		return helper.BuildErrorResponse("Invalid refresh token", "token mismatch", nil)
	}

	newAccessToken, err := auth.GenerateAccessToken(claims.UserID, claims.Username)
	if err != nil {
		return helper.BuildErrorResponse("Failed to generate access token", err.Error(), nil)
	}

	newRefreshToken, err := auth.GenerateRefreshToken(claims.UserID, claims.Username)
	if err != nil {
		return helper.BuildErrorResponse("Failed to generate refresh token", err.Error(), nil)
	}

	if err := s.saveRefreshToken(newRefreshToken); err != nil {
		return helper.BuildErrorResponse("Failed to save refresh token", err.Error(), nil)
	}

	if err := s.RefreshTokenRepo.DeleteByUserID(claims.UserID); err != nil {
		return helper.BuildErrorResponse("Failed to delete old refresh token", err.Error(), nil)
	}

	responseData := AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	return helper.BuildResponse(true, "Token refreshed successfully", responseData)
}

func NewAuthPortalService(portalRepo *repositories.PortalRepository, refreshTokenRepo *repositories.RefreshTokenRepository) AuthPortalServiceInterface {
	return &AuthPortalService{
		PortalRepo:       portalRepo,
		RefreshTokenRepo: refreshTokenRepo,
	}
}

type AuthPortalResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthPortalServiceInterface interface {
	Login(username, password string) helper.Response
	Logout(token string) helper.Response
	RefreshToken(token string) helper.Response
}
