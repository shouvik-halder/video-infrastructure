package helpers

import (
	"ApiGateway/configs"
	"ApiGateway/models"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.Claims{},
		func(token *jwt.Token) (any, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(configs.GetConfig().Auth.TOKENSECRET), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}