package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptor struct {
	Password string
}

func NewPasswordEncryptor() *PasswordEncryptor {
	return &PasswordEncryptor{}
}

func (encryptor *PasswordEncryptor) Encrypt() (string, error) {
	var passwordBytes = []byte(encryptor.Password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}

func (encryptor *PasswordEncryptor) IsHashedPasswordMatch(hashedPassword string) bool {
	byteHash := []byte(hashedPassword)
	byteOriginal := []byte(encryptor.Password)

	err := bcrypt.CompareHashAndPassword(byteHash, byteOriginal)
	if err != nil {
		fmt.Println("hashing matcher error : ", err)
		return false
	}

	return true
}
