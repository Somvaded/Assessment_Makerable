package repositories_test

import (
	"context"
	"database/sql"
	"regexp"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Somvaded/assessment/models"
	"github.com/Somvaded/assessment/repositories"
	"github.com/Somvaded/assessment/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindPatients_Success(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    ctx := context.Background()
    aadharID := "1234-5678-9012"
    rows := sqlmock.NewRows([]string{
        "id", "name", "phone", "age", "dob", "gender", "emergency_contact", "aadhar",
        "doctor_id", "payment_info", "known_allergies", "medications", "other_health_issues",
        "doctor_notes", "consent", "created_at", "updated_at",
    }).AddRow(1, "John Doe", "9876543210", 30, time.Now(), "male", "1234567890", aadharID,
        1, "Paid", "None", "Paracetamol", "None", "Healthy", true, time.Now(), time.Now())

    mock.ExpectQuery("SELECT \\* from patients where aadhar=\\$1;").WithArgs(aadharID).WillReturnRows(rows)

    patient, err := repositories.FindPatients(ctx, db, aadharID)
    assert.NoError(t, err)
    assert.Equal(t, "John Doe", patient.Name)
}

func TestFindPatients_NotFound(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    ctx := context.Background()
    mock.ExpectQuery("SELECT \\* from patients where aadhar=\\$1;").WillReturnError(sql.ErrNoRows)

    patient, err := repositories.FindPatients(ctx, db, "not-found")
    assert.Error(t, err)
    assert.Nil(t, patient)
}

func TestInsertPatient(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    ctx := context.Background()
    patient := models.Patient{
        Name: "John Doe", Phone: "9876543210", Age: 30, DOB: time.Now(),
        Gender: "male", EmergencyContact: "1234567890", Aadhar: "1234-5678-9012", DoctorID: 1,
        PaymentInfo: "Paid", KnownAllergies: "None", Medications: "Paracetamol",
        OtherHealthIssues: "None", DoctorNotes: "Healthy", Consent: true,
    }

    mock.ExpectQuery("INSERT INTO patients").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

    id, err := repositories.InsertPatient(ctx, db, patient)
    assert.NoError(t, err)
    assert.Equal(t, 1, id)
}

func TestUpdatePatient_Success(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    ctx := context.Background()
    patient := models.Patient{
        ID: 1, Name: "John Doe", Phone: "9876543210", Age: 30, DOB: time.Now(),
        Gender: "male", EmergencyContact: "1234567890", Aadhar: "1234-5678-9012", DoctorID: 1,
        PaymentInfo: "Paid", KnownAllergies: "None", Medications: "Paracetamol",
        OtherHealthIssues: "None", DoctorNotes: "Healthy", Consent: true,
    }

    mock.ExpectExec("UPDATE patients SET").WillReturnResult(sqlmock.NewResult(0, 1))
    mock.ExpectQuery("SELECT id, name, phone, age, dob, gender, emergency_contact").
        WithArgs(patient.ID).
        WillReturnRows(sqlmock.NewRows([]string{
            "id", "name", "phone", "age", "dob", "gender", "emergency_contact", "aadhar",
            "doctor_id", "payment_info", "known_allergies", "medications", "other_health_issues",
            "doctor_notes", "consent", "created_at", "updated_at",
        }).AddRow(
            1, "John Doe", "9876543210", 30, time.Now(), "male", "1234567890", "1234-5678-9012",
            1, "Paid", "None", "Paracetamol", "None", "Healthy", true, time.Now(), time.Now(),
        ))

    updatedPatient, err := repositories.UpdatePatient(ctx, db, patient)
    assert.NoError(t, err)
    assert.NotNil(t, updatedPatient)
}

func TestDeletePatient_Success(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    ctx := context.Background()
    patientID := 1
    mock.ExpectExec("DELETE FROM patients WHERE id = \\$1").
        WithArgs(patientID).WillReturnResult(sqlmock.NewResult(0, 1))

    err := repositories.DeletePatient(ctx, db, patientID)
    assert.NoError(t, err)
}

func TestDeletePatient_NotFound(t *testing.T) {
    db, mock, _ := sqlmock.New()
    defer db.Close()

    ctx := context.Background()
    patientID := 999
    mock.ExpectExec("DELETE FROM patients WHERE id = \\$1").
        WithArgs(patientID).WillReturnResult(sqlmock.NewResult(0, 0))

    err := repositories.DeletePatient(ctx, db, patientID)
    assert.Error(t, err)
}


func TestFindPatientsByDoctorID(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    ctx := context.Background()

    rows := sqlmock.NewRows([]string{
        "id", "name", "phone", "age", "gender", "emergency_contact",
        "known_allergies", "medications", "other_health_issues", 
        "doctor_notes", "consent", "created_at", "updated_at",
    }).AddRow(
        1, "John Doe", "1234567890", 30, "Male", "9876543210",
        "Peanuts", "Aspirin", "Asthma", "Note 1", true, time.Now(), time.Now(),
    )

    mock.ExpectQuery(regexp.QuoteMeta(`
        SELECT id , name , phone , age,gender, emergency_contact,
        known_allergies, medications, other_health_issues, 
        doctor_notes,consent, created_at, updated_at FROM patients
        WHERE doctor_id = $1;
    `)).
        WithArgs(1).
        WillReturnRows(rows)

    patients, err := repositories.FindPatientsByDoctorID(ctx, db, 1)
    assert.NoError(t, err)
    assert.Len(t, patients, 1)
    assert.Equal(t, "John Doe", patients[0].Name)
}

func TestUpdateMedicalInfo(t *testing.T) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    ctx := context.Background()
    now := time.Now()

    updateInfo := models.DocPatientUpdate{
        KnownAllergies:   "Dust",
        Medications:      "Paracetamol",
        OtherHealthIssues: "None",
        DoctorNotes:      "Stable condition",
    }

    row := sqlmock.NewRows([]string{
        "id", "name", "phone", "age", "gender", "emergency_contact",
        "known_allergies", "medications", "other_health_issues", 
        "doctor_notes", "consent", "created_at", "updated_at",
    }).AddRow(
        1, "Jane Doe", "9876543210", 28, "Female", "1234567890",
        "Dust", "Paracetamol", "None", "Stable condition", true, now, now,
    )

    mock.ExpectQuery(regexp.QuoteMeta(`
        UPDATE patients
        SET known_allergies = $1, medications = $2, other_health_issues = $3, doctor_notes = $4, updated_at = NOW()
        WHERE id = $5
        RETURNING id, name, phone, age, gender, emergency_contact, known_allergies, medications, other_health_issues, doctor_notes, consent, created_at, updated_at;
    `)).
        WithArgs(updateInfo.KnownAllergies, updateInfo.Medications, updateInfo.OtherHealthIssues, updateInfo.DoctorNotes, 1).
        WillReturnRows(row)

    updatedPatient, err := repositories.UpdateMedicalInfo(ctx, db, 1, updateInfo)
    assert.NoError(t, err)
    assert.Equal(t, "Jane Doe", updatedPatient.Name)
    assert.Equal(t, "Dust", updatedPatient.KnownAllergies)
}

func TestFindUserByEmail_Doctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	password := "securepass"
	hash, _ := utils.HashPassword(password)

	rows := sqlmock.NewRows([]string{
		"id", "email", "role", "password_hash",
		"doctor_name", "specialty", "emergency_contact", "license_number", "experience_years", "d_created_at", "d_updated_at",
		"receptionist_name", "phone", "r_created_at", "r_updated_at",
	}).AddRow(
		1, "doc@example.com", "doctor", hash,
		"Dr. Smith", "Cardiology", "1234567890", "LIC1234", 10, time.Now(), time.Now(),
		nil, nil, nil, nil,
	)

	mock.ExpectQuery("SELECT (.+) FROM users u").
		WithArgs("doc@example.com").
		WillReturnRows(rows)

	user, doctor, receptionist, err := repositories.FindUserByEmail(context.Background(), db, "doc@example.com", password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, doctor)
	assert.Nil(t, receptionist)
	assert.Equal(t, "doctor", user.Role)
	assert.Equal(t, "Dr. Smith", doctor.Name)
}

func TestFindUserByEmail_Receptionist(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	password := "adminpass"
	hash, _ := utils.HashPassword(password)

	rows := sqlmock.NewRows([]string{
		"id", "email", "role", "password_hash",
		"doctor_name", "specialty", "emergency_contact", "license_number", "experience_years", "d_created_at", "d_updated_at",
		"receptionist_name", "phone", "r_created_at", "r_updated_at",
	}).AddRow(
		2, "rec@example.com", "receptionist", hash,
		nil, nil, nil, nil, nil, nil, nil,
		"Alice", "9876543210", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT (.+) FROM users u").
		WithArgs("rec@example.com").
		WillReturnRows(rows)

	user, doctor, receptionist, err := repositories.FindUserByEmail(context.Background(), db, "rec@example.com", password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Nil(t, doctor)
	assert.NotNil(t, receptionist)
	assert.Equal(t, "receptionist", user.Role)
	assert.Equal(t, "Alice", receptionist.Name)
}

func TestFindUserByEmail_WrongPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	hash, _ := utils.HashPassword("correctpass")

	rows := sqlmock.NewRows([]string{
		"id", "email", "role", "password_hash",
		"doctor_name", "specialty", "emergency_contact", "license_number", "experience_years", "d_created_at", "d_updated_at",
		"receptionist_name", "phone", "r_created_at", "r_updated_at",
	}).AddRow(
		3, "user@example.com", "doctor", hash,
		"Dr. Wrong", "Neuro", "0001112222", "LIC5678", 5, time.Now(), time.Now(),
		nil, nil, nil, nil,
	)

	mock.ExpectQuery("SELECT (.+) FROM users u").
		WithArgs("user@example.com").
		WillReturnRows(rows)

	user, doctor, receptionist, err := repositories.FindUserByEmail(context.Background(), db, "user@example.com", "wrongpass")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Nil(t, doctor)
	assert.Nil(t, receptionist)
	assert.Contains(t, err.Error(), "error comparing password")
}

func TestFindUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM users u").
		WithArgs("missing@example.com").
		WillReturnError(sql.ErrNoRows)

	user, doctor, receptionist, err := repositories.FindUserByEmail(context.Background(), db, "missing@example.com", "any")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Nil(t, doctor)
	assert.Nil(t, receptionist)
	assert.Contains(t, err.Error(), "no user found")
}