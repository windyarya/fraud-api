package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	UserGroupID uint `json:"user_group_id"`
	WorkGroupID uint `json:"work_group_id"`
	jwt.RegisteredClaims
}

var SigningKey = []byte(os.Getenv("TOKEN_SALT"))

func ClaimToken(id uint, userGroupID uint, workGroupID uint) (string, error) {
	claims := &JwtCustomClaims{
		id,
		userGroupID,
		workGroupID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(t string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(t, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}