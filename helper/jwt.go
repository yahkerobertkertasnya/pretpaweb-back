package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateJWT(userId string) (string, error) {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(GoDotEnvVariable("JWT_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {

	secretKey := []byte(GoDotEnvVariable("JWT_KEY"))

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiration := time.Unix(int64(claims["exp"].(float64)), 0)

		if time.Now().After(expiration) {
			return nil, fmt.Errorf("token expired")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
