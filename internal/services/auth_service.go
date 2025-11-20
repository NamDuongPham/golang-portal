package services

import (
	"errors"
	"log"

	"github.com/namduong/project-layout/internal/auth"
	"github.com/namduong/project-layout/internal/models"
	"github.com/namduong/project-layout/internal/repositories"
	"github.com/namduong/project-layout/utils"
)

type AuthService struct {
	AdminRepo        *repositories.AdminRepository
	RefreshTokenRepo *repositories.RefreshTokenRepository
}

func (s *AuthService) Login(username string, password string) (accessToken string, refreshToken string, err error) {
	admin, err := s.AdminRepo.FindByUsername(username)
	if err != nil {
		return "", "", err
	}
	hashedPassWord := utils.HashPassword(password)
	if admin.Password != hashedPassWord {
		return "", "", errors.New("wrong password")
	}

	accessToken, err = auth.GenerateAccessToken(admin.ID, admin.UserName)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = auth.GenerateRefreshToken(admin.ID, admin.UserName)
	if err != nil {
		return "", "", err
	}

	claims, err := auth.DecodeRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	rt := &models.RefreshToken{
		ID:        claims.Id,
		Token:     refreshToken,
		UserID:    claims.UserID,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	err = s.RefreshTokenRepo.Create(rt)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
func (s *AuthService) Logout(token string) error {
	claims, err := auth.DecodeAccessToken(token)
	if err != nil {
		return err
	}
	log.Printf("Logging out user ID: %s", claims.UserID)
	err = s.RefreshTokenRepo.DeleteByUserID(claims.UserID)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthService(adminRepo *repositories.AdminRepository, refreshTokenRepo *repositories.RefreshTokenRepository) AuthServiceInterface {
	return &AuthService{
		AdminRepo:        adminRepo,
		RefreshTokenRepo: refreshTokenRepo,
	}
}

type AuthServiceInterface interface {
	Login(username, password string) (accessToken, refreshToken string, err error)
	Logout(token string) error
}
