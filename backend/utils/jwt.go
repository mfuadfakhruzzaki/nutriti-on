// backend/utils/jwt.go
package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTUtil struct {
    Secret string
}

type Claims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

func NewJWTUtil(secret string) *JWTUtil {
    return &JWTUtil{Secret: secret}
}

func (j *JWTUtil) GenerateToken(userID uint, email string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // Token valid for 3 days
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.Secret))
}

func (j *JWTUtil) ValidateToken(tokenStr string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(j.Secret), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
