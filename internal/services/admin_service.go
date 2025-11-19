package services

import (
	"errors"

	"github.com/namduong/project-layout/internal/auth"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
	"github.com/namduong/project-layout/utils"
)

type AuthService interface {
	Login(username, password string) (accessToken, refreshToken string, err error)
}

type authService struct {
	AdminRepo    repositories.AdminRepository
	RefreshToken repositories.RefreshTokenRepository
	TokenService auth.TokenService
}

func (s *authService) Login(username, password string) (accessToken, refreshToken string, err error) {
	admin, err := s.AdminRepo.FindByUsername(username)
	if err != nil {
		return "", "", err
	}
	hashedPassWord := utils.HashPassword(password)
	if admin.Password != hashedPassWord {
		return "", "", errors.New("wrong password")
	}
	accessToken, err = s.TokenService.GenerateAccessToken(admin.ID, admin.UserName)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.TokenService.GenerateRefreshToken(admin.ID, admin.UserName)
	if err != nil {
		return "", "", err
	}
	claims, err := s.TokenService.DecodeRefreshToken(refreshToken)

	rt := &models.RefreshToken{
		Token:     refreshToken,
		UserID:    claims.UserID,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	err = s.RefreshToken.Create(rt)
	if err != nil {
		return "", "", err
	}

	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
