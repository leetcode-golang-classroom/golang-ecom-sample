package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	var JWTExpirationInSeconds int64 = 3600
	if config.C != nil {
		JWTExpirationInSeconds = config.C.JWTExpirationInSeconds
	}
	expiration := time.Second * time.Duration(JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
