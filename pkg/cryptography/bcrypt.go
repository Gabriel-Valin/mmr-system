package cryptography

import "golang.org/x/crypto/bcrypt"

func HashPassword(plaintext string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
