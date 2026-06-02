package utils

import (
	"errors"
	"time"

	"example.com/models"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId int64  `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var signToken = "random token"

func SignJwt(payload models.User) string {
	claims := CustomClaims{
		UserId: int64(payload.Id),
		Email:  payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	tokenBuilder := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := tokenBuilder.SignedString([]byte(signToken))

	if err != nil {
		panic(err)
	}

	return jwt
}

func ParseJwt(token string) (*CustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signToken), nil
	})

	if err != nil {
		return nil, err
	}

	jti, err := parsedToken.Claims.GetExpirationTime()

	if err != nil {
		return nil, err
	}

	ttl := time.Until(jti.Time)

	if ttl <= 0 {
		return nil, errors.New("Token is expired")
	}

	return parsedToken.Claims.(*CustomClaims), nil
}
