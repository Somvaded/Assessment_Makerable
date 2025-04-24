package models

import "time"

type Receptionist struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

