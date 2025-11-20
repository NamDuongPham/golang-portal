package services

// func (s *AuthService) Login(username, password string) (accessToken, refreshToken string, err error) {
// 	admin, err := s.AdminRepo.FindByUsername(username)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	hashedPassWord := utils.HashPassword(password)
// 	if admin.Password != hashedPassWord {
// 		return "", "", errors.New("wrong password")
// 	}

// 	accessToken, err = auth.GenerateAccessToken(admin.ID, admin.UserName)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	refreshToken, err = auth.GenerateRefreshToken(admin.ID, admin.UserName)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	claims, err := auth.DecodeRefreshToken(refreshToken)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	rt := &models.RefreshToken{
// 		Token:     refreshToken,
// 		UserID:    claims.UserID,
// 		ExpiresAt: claims.ExpiresAt.Time,
// 	}

// 	err = s.RefreshTokenRepo.Create(rt)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return accessToken, refreshToken, nil
// }
