package helper

import (
	"dennydolok/BSP-BE/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(role, id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["role"] = role
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Secret))
}

func GetClaimsRole(reqToken string) float64 {
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 1 {
		return 0
	}
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return 0
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["role"].(float64)
}

func GetClaimsID(reqToken string) int {
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 1 {
		return 0
	}
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return 0
	}
	claims := token.Claims.(jwt.MapClaims)
	if userID, ok := claims["user_id"].(float64); ok {
		return int(userID)
	}
	return 0
}
