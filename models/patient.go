package models

import "time"

type Patient struct {
	ID                int       `json:"id,omitempty"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Age               int       `json:"age"`
	DOB               time.Time `json:"dob"`
	Gender            string    `json:"gender"`
	EmergencyContact  string    `json:"emergency_contact"`
	Aadhar            string    `json:"aadhar"`
	DoctorID          int       `json:"doctor_id"`
	PaymentInfo       string    `json:"payment_info"`
	KnownAllergies    string    `json:"known_allergies"`
	Medications       string    `json:"medications"`
	OtherHealthIssues string    `json:"other_health_issues"`
	DoctorNotes       string    `json:"doctor_notes"`
	Consent           bool      `json:"consent"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}