package services

import (
	"github.com/namduong/project-layout/helper"
	"github.com/namduong/project-layout/internal/auth"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AdminRepo        *repositories.AdminRepository
	RefreshTokenRepo *repositories.RefreshTokenRepository
}

func (s *AuthService) saveRefreshToken(rawToken string) error {
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

func (s *AuthService) Login(username string, password string) helper.Response {
	admin, err := s.AdminRepo.FindByUsername(username)
	if err != nil {
		return helper.BuildErrorResponse("User not found", "invalid credentials", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return helper.BuildErrorResponse("Wrong password", "invalid credentials", nil)
	}

	accessToken, err := auth.GenerateAccessToken(admin.ID, admin.UserName)
	if err != nil {
		return helper.BuildErrorResponse("Failed to generate access token", err.Error(), nil)
	}

	refreshToken, err := auth.GenerateRefreshToken(admin.ID, admin.UserName)
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

func (s *AuthService) Logout(token string) helper.Response {
	claims, err := auth.DecodeAccessToken(token)
	if err != nil {
		return helper.BuildErrorResponse("Invalid access token", err.Error(), nil)
	}

	if err := s.RefreshTokenRepo.DeleteByUserID(claims.UserID); err != nil {
		return helper.BuildErrorResponse("Failed to logout", err.Error(), nil)
	}

	return helper.BuildResponse(true, "Logout successful", nil)
}

func (s *AuthService) RefreshToken(token string) helper.Response {
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

func NewAuthService(adminRepo *repositories.AdminRepository, refreshTokenRepo *repositories.RefreshTokenRepository) AuthServiceInterface {
	return &AuthService{
		AdminRepo:        adminRepo,
		RefreshTokenRepo: refreshTokenRepo,
	}
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthServiceInterface interface {
	Login(username, password string) helper.Response
	Logout(token string) helper.Response
	RefreshToken(token string) helper.Response
}
