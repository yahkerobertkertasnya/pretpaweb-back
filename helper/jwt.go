package helper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateJWT(user string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(GoDotEnvVariable("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
