package infrastructure

import (
	"task_manager_refactored/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)


var jwtKey = []byte("1234")

// GenerateJWT generates a JWT token for the given user ID, username, and role.
// The token expires in 24 hours.
func GenerateJWT(userID string, username string, role string) (string, error) {
	claims := &domain.Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}