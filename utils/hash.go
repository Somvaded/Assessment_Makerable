package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}


func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Error comparing password:", err)
	}
	return err
}
