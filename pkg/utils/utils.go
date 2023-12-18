package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the input password and returns the hashed password.
func HashPassword(password string) *string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	result := string(hashedPassword)

	return &result
}

// VerifyPassword checks if the input password matches the hashed password.
func VerifyPassword(inputPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
