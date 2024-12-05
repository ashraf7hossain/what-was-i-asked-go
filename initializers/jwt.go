package initializers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TODO: change secret key

var SecretKey = os.Getenv("JWT_SECRET")

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp"   : time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
