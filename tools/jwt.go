package tools

import (
	"context"
	"errors"
	"nozzlium/kepo_backend/constants"
	"nozzlium/kepo_backend/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"userId"`
}

func GetClaimsFromContext(
	context context.Context,
) (JwtClaims, error) {
	claims, ok := context.Value(constants.USER_ID_CLAIMS).(JwtClaims)
	if !ok {
		return claims, exception.UnauthorizedError{}
	}
	return claims, nil
}

func NewJwtToken(userId uint) (string, error) {
	signKey := []byte(constants.SIGNATURE_KEY)
	claims := JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		UserId: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(signKey)
	return signed, err
}

func ParseJWTToken(tokenString string) (*JwtClaims, error) {
	authClaims := JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &authClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.SIGNATURE_KEY), nil
	})
	if err != nil {
		return &authClaims, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return &authClaims, errors.New("INVALID TOKEN")
}
