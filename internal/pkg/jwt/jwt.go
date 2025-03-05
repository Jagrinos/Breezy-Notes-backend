package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"uasbreezy/config"
)

func GenerateToken(login string, TYPE string) (string, error) {
	//var claims claimsLogin
	var claims jwt.Claims
	switch TYPE {
	case "ACCESS":
		claims = jwt.MapClaims{
			"sub": login,
			"exp": time.Now().Add(time.Minute * 10).Unix(),
		}
	case "REFRESH":
		claims = jwt.MapClaims{
			"sub": login,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		}
	default:
		return "", errors.New("undefined type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(config.PRIVATEKEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Refresh(refreshToken string, login string) (string, error) {
	_, err := VerifyToken(refreshToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", errors.New("refresh token is expired")
		}
		return "", err
	}
	return GenerateToken(login, "ACCESS")
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.PUBLICKEY, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetLoginFromToken(tokenString string) (string, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return "", err
	}
	login, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}
	return login, nil
}
