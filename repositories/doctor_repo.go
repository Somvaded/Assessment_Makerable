package repositories


import (	
"context"
"database/sql"
"fmt"
"github.com/Somvaded/assessment/models"
)

func FindPatientsByDoctorID(ctx context.Context, db *sql.DB, doctorID int) ([]models.DocPatientResponse, error) {
	var patients []models.DocPatientResponse

	query := `
	SELECT id , name , phone , age,gender, emergency_contact,
	known_allergies, medications, other_health_issues, 
	doctor_notes,consent, created_at, updated_at FROM patients
	WHERE doctor_id = $1;
	`

	rows, err := db.QueryContext(ctx, query, doctorID)
	if err != nil {
		return nil, fmt.Errorf("error querying patients: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var patient models.DocPatientResponse
		err := rows.Scan(
			&patient.ID,
			&patient.Name,
			&patient.Phone,
			&patient.Age,
			&patient.Gender,
			&patient.EmergencyContact,
			&patient.KnownAllergies,
			&patient.Medications,
			&patient.OtherHealthIssues,
			&patient.DoctorNotes,
			&patient.Consent,
			&patient.CreatedAt,
			&patient.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning patient: %w", err)
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return patients, nil
}


func UpdateMedicalInfo(ctx context.Context, db *sql.DB, patient_id int,updateInfo models.DocPatientUpdate)(*models.DocPatientResponse,error){
	query := `
	UPDATE patients
	SET known_allergies = $1, medications = $2, other_health_issues = $3, doctor_notes = $4, updated_at = NOW()
	WHERE id = $5
	RETURNING id, name, phone, age, gender, emergency_contact, known_allergies, medications, other_health_issues, doctor_notes, consent, created_at, updated_at;
	`

	var updatedPatient models.DocPatientResponse
	err := db.QueryRowContext(
		ctx,
		query,
		updateInfo.KnownAllergies,
		updateInfo.Medications,
		updateInfo.OtherHealthIssues,
		updateInfo.DoctorNotes,
		patient_id,
	).Scan(
		&updatedPatient.ID,
		&updatedPatient.Name,
		&updatedPatient.Phone,
		&updatedPatient.Age,
		&updatedPatient.Gender,
		&updatedPatient.EmergencyContact,
		&updatedPatient.KnownAllergies,
		&updatedPatient.Medications,
		&updatedPatient.OtherHealthIssues,
		&updatedPatient.DoctorNotes,
		&updatedPatient.Consent,
		&updatedPatient.CreatedAt,
		&updatedPatient.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating patient medical info: %w", err)
	}

	return &updatedPatient, nil
}

