package models

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	PasswordHash string `json:"password_hash"`
}