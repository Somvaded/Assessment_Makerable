package models

import "time"

type Doctor struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Specialty        string    `json:"specialty"`
	EmergencyContact string    `json:"emergency_contact"`
	LicenseNumber    string    `json:"license_number"`
	ExperienceYears  int       `json:"experience_years"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

type DocPatientResponse struct {
	ID                int       `json:"id,omitempty"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Age               int       `json:"age"`
	Gender            string    `json:"gender"`
	EmergencyContact  string    `json:"emergency_contact"`
	KnownAllergies    string    `json:"known_allergies"`
	Medications       string    `json:"medications"`
	OtherHealthIssues string    `json:"other_health_issues"`
	DoctorNotes       string    `json:"doctor_notes"`
	Consent           bool      `json:"consent"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}


type DocPatientUpdate struct {
	DoctorNotes       string    `json:"doctor_notes"`
	Medications       string    `json:"medications"`
	OtherHealthIssues string    `json:"other_health_issues"`
	KnownAllergies 	  string    `json:"known_allergies"`
}