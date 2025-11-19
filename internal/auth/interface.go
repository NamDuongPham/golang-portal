package auth

type TokenService interface {
	GenerateAccessToken(userID, username string) (string, error)
	GenerateRefreshToken(userID, username string) (string, error)
	DecodeAccessToken(tokenString string) (*Claims, error)
	DecodeRefreshToken(tokenString string) (*Claims, error)
}
