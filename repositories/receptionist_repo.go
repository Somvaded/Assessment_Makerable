package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/Somvaded/assessment/models"
)





func FindPatients(ctx context.Context,DB *sql.DB, aadharid string) (*models.Patient,error){
	var patient models.Patient

	query := `
	SELECT * from patients
	where aadhar=$1;
	`
	err := DB.QueryRowContext(ctx,query,aadharid).Scan(
		&patient.ID,
		&patient.Name,
		&patient.Phone,
		&patient.Age,
		&patient.DOB,
		&patient.Gender,
		&patient.EmergencyContact,
		&patient.Aadhar,
		&patient.DoctorID,
		&patient.PaymentInfo,
		&patient.KnownAllergies,
		&patient.Medications,
		&patient.OtherHealthIssues,
		&patient.DoctorNotes,
		&patient.Consent,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil ,fmt.Errorf("no patient found")
		}else {
			return nil , fmt.Errorf(err.Error())
		}
	}
	return &patient,nil
}

func InsertPatient(ctx context.Context,db *sql.DB, patient models.Patient)(int, error) {
	parsedDob, err := time.Parse("2006-01-02", patient.DOB.Format("2006-01-02"))
    if err != nil {
        return 0, fmt.Errorf("error parsing date: %w", err)
    }
	query := `
	INSERT INTO patients (
		name, phone, age, dob, gender,
		emergency_contact, aadhar, doctor_id,
		payment_info, known_allergies, medications, other_health_issues,
		doctor_notes, consent
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7,
		$8, $9, $10,
		$11, $12, $13, $14
	)
	RETURNING id;`

	var id int
	err = db.QueryRowContext(
		ctx,
		query,
		patient.Name,
		patient.Phone,
		patient.Age,
		parsedDob,
		patient.Gender,
		patient.EmergencyContact,
		patient.Aadhar,
		patient.DoctorID,
		patient.PaymentInfo,
		patient.KnownAllergies,
		patient.Medications,
		patient.OtherHealthIssues,
		patient.DoctorNotes,
		patient.Consent,
	).Scan(&id)
	if err != nil {
		return 0,fmt.Errorf("query error %w",err)
	}
	return id,nil
}

func UpdatePatient(ctx context.Context,db *sql.DB, patient models.Patient) (*models.Patient, error) {
	// Parse the date of birth
    parsedDob, err := time.Parse("2006-01-02", patient.DOB.Format("2006-01-02"))
    if err != nil {
        return nil, fmt.Errorf("error parsing date: %w", err)
    }

    // Query to update patient details
    query := `
    UPDATE patients SET
        name = $1, phone = $2, age = $3, dob = $4, gender = $5,
        emergency_contact = $6,aadhar = $7, doctor_id = $8, payment_info = $9,
        known_allergies = $10, medications = $11, other_health_issues = $12,
        doctor_notes = $13, consent = $14, updated_at = NOW()
    WHERE id = $15;
    `

    row,err := db.ExecContext(
		ctx,
        query,
        patient.Name,
        patient.Phone,
        patient.Age,
        parsedDob,
        patient.Gender,
        patient.EmergencyContact,
		patient.Aadhar,
        patient.DoctorID,
        patient.PaymentInfo,
        patient.KnownAllergies,
        patient.Medications,
        patient.OtherHealthIssues,
        patient.DoctorNotes,
        patient.Consent,
        patient.ID,
    )
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	rowsAffected, err := row.RowsAffected()
    if err != nil {
        return nil, fmt.Errorf("error checking rows affected: %w", err)
    }
    if rowsAffected == 0 {
        return nil, fmt.Errorf("no patient found with ID %d", patient.ID)
    }

	 selectQuery := `
	 SELECT 
		 id, name, phone, age, dob, gender, emergency_contact, aadhar , doctor_id,
		 payment_info, known_allergies, medications, other_health_issues,
		 doctor_notes, consent, created_at, updated_at
	 FROM patients
	 WHERE id = $1;
	 `
 
	 var updatedPatient models.Patient
	 err = db.QueryRowContext(ctx, selectQuery, patient.ID).Scan(
		 &updatedPatient.ID,
		 &updatedPatient.Name,
		 &updatedPatient.Phone,
		 &updatedPatient.Age,
		 &updatedPatient.DOB,
		 &updatedPatient.Gender,
		 &updatedPatient.EmergencyContact,
		 &updatedPatient.Aadhar,
		 &updatedPatient.DoctorID,
		 &updatedPatient.PaymentInfo,
		 &updatedPatient.KnownAllergies,
		 &updatedPatient.Medications,
		 &updatedPatient.OtherHealthIssues,
		 &updatedPatient.DoctorNotes,
		 &updatedPatient.Consent,
		 &updatedPatient.CreatedAt,
		 &updatedPatient.UpdatedAt,
	 )
	 if err != nil {
		 return nil, fmt.Errorf("fetch updated patient error: %w", err)
	 }
 
	 return &updatedPatient, nil

}

func DeletePatient(ctx context.Context,db *sql.DB, patientid int) (error) {
	query := `
	DELETE FROM patients WHERE id = $1;
	`
	result, err := db.ExecContext(ctx, query, patientid)
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no patient found with id %d", patientid)
	}
	return nil
}
