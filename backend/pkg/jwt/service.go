package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT claims structure
type Claims struct {
	UserID      string   `json:"user_id"`
	SchoolID    *string  `json:"school_id,omitempty"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions,omitempty"`
	jwt.RegisteredClaims
}

// Service handles JWT token generation and validation
type Service struct {
	secretKey     []byte
	accessExpiry  time.Duration
	refreshExpiry time.Duration
	issuer        string
}

// NewService creates a new JWT service
func NewService(secretKey string, accessExpiry, refreshExpiry time.Duration, issuer string) *Service {
	return &Service{
		secretKey:     []byte(secretKey),
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
		issuer:        issuer,
	}
}

// GenerateAccessToken generates a new JWT access token
func (s *Service) GenerateAccessToken(userID, role string, schoolID *string, permissions []string) (string, error) {
	claims := &Claims{
		UserID:      userID,
		SchoolID:    schoolID,
		Role:        role,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

// GenerateRefreshToken generates a new JWT refresh token
func (s *Service) GenerateRefreshToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

// ValidateToken validates a JWT token and returns the claims
func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ExtractUserID extracts user ID from token without full validation (for refresh token)
func (s *Service) ExtractUserID(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}
