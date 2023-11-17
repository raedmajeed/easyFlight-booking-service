package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPass, nil
}

func CheckPasswordMatch(hashedPass []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	return err == nil
}
