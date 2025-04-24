package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Somvaded/assessment/models"
	"github.com/Somvaded/assessment/utils"
)

func FindUserByEmail(ctx context.Context,db *sql.DB, email string,password string) (*models.User, *models.Doctor, *models.Receptionist, error) {
	
	query := `
	SELECT 
	u.id, u.email, u.role,u.password_hash,
		d.name, d.specialty, d.emergency_contact, d.license_number, d.experience_years,d.created_at, d.updated_at,
		r.name, r.phone,r.created_at, r.updated_at
	FROM users u
	LEFT JOIN doctors d ON u.email = d.email
	LEFT JOIN receptionists r ON u.email = r.email
	WHERE u.email = $1;
	`

	row := db.QueryRowContext(ctx, query, email)

	var (
		user                      models.User
		nullDoctorName           sql.NullString
		nullSpecialty            sql.NullString
		nullEmergencyContactInfo sql.NullString
		nullLicense              sql.NullString
		nullExperience           sql.NullInt32
		nullReceptionistName     sql.NullString
		nullReceptionistPhone    sql.NullString
		nullRCreatedAt           sql.NullTime
		nullRUpdatedAt           sql.NullTime
		nullDCreatedAt           sql.NullTime
		nullDUpdatedAt           sql.NullTime

	)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Role,
		&user.PasswordHash,

		&nullDoctorName,
		&nullSpecialty,
		&nullEmergencyContactInfo,
		&nullLicense,
		&nullExperience,
		&nullDCreatedAt,
		&nullDUpdatedAt,

		&nullReceptionistName,
		&nullReceptionistPhone,
		&nullRCreatedAt,
		&nullRUpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil, fmt.Errorf("no user found %w",err)
		}
		return nil, nil, nil, err
	}

	err = utils.ComparePassword(user.PasswordHash, password)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error comparing password %w", err)
	}
	var doctorProfile *models.Doctor
	if user.Role == "doctor" {
		doctorProfile = &models.Doctor{
			ID:              user.ID,
			Name:            nullDoctorName.String,
			Email:           user.Email,
			Specialty:       nullSpecialty.String,
			EmergencyContact: nullEmergencyContactInfo.String,
			LicenseNumber:   nullLicense.String,
			ExperienceYears: int(nullExperience.Int32),
			CreatedAt: 	 nullDCreatedAt.Time,
			UpdatedAt: 	 nullDUpdatedAt.Time,
		}
	}

	var receptionistProfile *models.Receptionist
	if user.Role == "receptionist" {
		receptionistProfile = &models.Receptionist{
			ID:    user.ID,
			Name:  nullReceptionistName.String,
			Email: user.Email,
			Phone: nullReceptionistPhone.String,
			CreatedAt: nullRCreatedAt.Time,
			UpdatedAt: nullRUpdatedAt.Time,
		}
	}

	return &user, doctorProfile, receptionistProfile, nil
}


