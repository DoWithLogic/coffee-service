package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// UserClaims represents the claims in the JWT token
type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Authorization string

// GenerateToken generates a JWT token for a given username and password
func (c *Authorization) GenerateToken(email string) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, including the username and expiration time
	claims := &UserClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token using the claims and sign it with a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(*c))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
