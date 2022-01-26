package jwt

import (
	"course_management/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenJWT struct {
	Email string
}

type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

func (t *TokenJWT) GenerateToken() (string, error) {
	duration := time.Duration(config.EnvConfig.JWT_LIFE_MINUTES) * time.Minute
	expiresAt := time.Now().Add(duration).Unix()
	claims := &Claims{
		UserID:         t.Email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiresAt},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.EnvConfig.JWT_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *TokenJWT) DecodeToken(jwtToken string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		jwtToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return config.EnvConfig.JWT_KEY, nil
		},
	)

	if err != nil || !token.Valid {
		return "", err
	}

	return claims.UserID, nil
}
