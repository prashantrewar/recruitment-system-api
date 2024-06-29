package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"recruitment-system/models"
)

func GenerateJWT(user models.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"userID":    user.ID,
		"userType":  user.UserType,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateJWT(tokenStr string, secretKey string) (*models.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	user := &models.User{
		ID:       uint(claims["userID"].(float64)), // Ensure conversion is correct
		UserType: claims["userType"].(string),
		Email:    claims["email"].(string),
	}

	return user, nil
}
