package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your-secret-key") // Replace with a strong secret key

// TokenPair holds access and refresh tokens
type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// GenerateTokenPair generates a new pair of access and refresh tokens
func GenerateTokenPair(userID uint) (*TokenPair, error) {
	// Access token
	expirationTime := time.Now().Add(15 * time.Minute) // Adjust as needed
	accessClaims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "your-app-name",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Refresh token
	refreshExpirationTime := time.Now().Add(7 * 24 * time.Hour) // Adjust as needed
	refreshClaims := &jwt.StandardClaims{
		ExpiresAt: refreshExpirationTime.Unix(),
		Issuer:    "your-app-name",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}, nil
}

// VerifyAccessToken verifies the access token
func VerifyAccessToken(tokenString string) (*jwt.StandardClaims, error) {
	return verifyToken(tokenString)
}

// VerifyRefreshToken verifies the refresh token
func VerifyRefreshToken(tokenString string) (*jwt.StandardClaims, error) {
	return verifyToken(tokenString)
}

// Helper function to verify a token
func verifyToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
