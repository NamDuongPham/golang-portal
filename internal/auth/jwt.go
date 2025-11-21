package auth

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	config "github.com/namduong/project-layout/configs"
)

type Claims struct {
	Id       string `json:"id,omitempty"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func getAccessSecretKey() []byte {
	secret := config.Cfg.JWT.AccessSecretKey
	if secret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET is not set in environment variables")
	}
	return []byte(secret)
}
func getRefreshSecretKey() []byte {
	secret := config.Cfg.JWT.RefreshSecretKey
	if secret == "" {
		log.Fatal("REFRESH_TOKEN_SECRET is not set in environment variables")
	}
	return []byte(secret)
}
func GenerateAccessToken(userID, username string) (string, error) {
	ttl := time.Minute * 15
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getAccessSecretKey())
}
func DecodeAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getAccessSecretKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid or expired token")
}
func GenerateRefreshToken(userID, username string) (string, error) {
	ttl := time.Hour * 24
	tokenID := uuid.New().String()
	claims := Claims{
		Id:       tokenID,
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getRefreshSecretKey())
}
func DecodeRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getRefreshSecretKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid or expired token")
}

type TokenService interface {
	GenerateAccessToken(userID, username string) (string, error)
	GenerateRefreshToken(userID, username string) (string, error)
	DecodeAccessToken(tokenString string) (*Claims, error)
	DecodeRefreshToken(tokenString string) (*Claims, error)
}
