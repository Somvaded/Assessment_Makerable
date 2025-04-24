package utils_test

import (
	"testing"
	"time"

	"github.com/Somvaded/assessment/models"
	"github.com/Somvaded/assessment/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	password := "myStrongP@ssw0rd"

	hashed, err := utils.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	assert.NotEqual(t, password, hashed, "Hashed password should not match original password")
}

func TestComparePassword_Success(t *testing.T) {
	password := "myStrongP@ssw0rd"
	hashed, err := utils.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = utils.ComparePassword(hashed, password)
	assert.NoError(t, err, "Password should match hashed password")
}

func TestComparePassword_Failure(t *testing.T) {
	password := "correctPassword"
	wrongPassword := "wrongPassword"

	hashed, err := utils.HashPassword(password)
	assert.NoError(t, err)

	err = utils.ComparePassword(hashed, wrongPassword)
	assert.Error(t, err, "Comparison should fail with incorrect password")
}

func TestGenerateAndVerifyJWT(t *testing.T) {
	// Set a mock secret
	userID := 42
	role := "doctor"

	tokenStr, err := utils.GenerateJWT(userID, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	claims, err := utils.VerifyJWT(tokenStr)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, role, claims.Role)
	assert.WithinDuration(t, time.Now().Add(24*time.Hour), claims.ExpiresAt.Time, time.Minute)
	assert.WithinDuration(t, time.Now(), claims.IssuedAt.Time, time.Minute)
}

func TestVerifyJWT_InvalidToken(t *testing.T) {

	_, err := utils.VerifyJWT("this.is.not.a.valid.token")
	assert.Error(t, err)
}

func TestVerifyJWT_TamperedToken(t *testing.T) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		UserID: 99,
		Role:   "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})

	tamperedToken, _ := token.SignedString([]byte("wrongsecret"))

	_, err := utils.VerifyJWT(tamperedToken)
	assert.Error(t, err)
}
