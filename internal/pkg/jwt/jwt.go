package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"uasbreezy/config"
)

func GenerateToken(login string, TYPE string) (string, error) {
	var timeLife int64

	switch TYPE {
	case "ACCESS":
		timeLife = time.Now().Add(time.Minute * 10).Unix()
	case "REFRESH":
		timeLife = time.Now().Add(time.Hour * 24 * 7).Unix()
	default:
		return "", errors.New("undefined type")
	}

	claims := jwt.MapClaims{
		"login": login,
		"exp":   timeLife,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(config.SECRETKEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Refresh(refreshToken string) (string, error) {
	token, err := verifyToken(refreshToken)
	if err != nil {
		return "", errors.New("refresh token is expired")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	login := claims["login"].(string)

	return GenerateToken(login, "ACCESS")
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.SECRETKEY, nil
	})

	return token, err
}
