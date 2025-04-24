package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/Somvaded/assessment/models"
	"github.com/Somvaded/assessment/repositories"
	"github.com/gin-gonic/gin"
)

type ReceptionistHandler struct {
	DB *sql.DB
}

func NewReceptionistHandler(db *sql.DB) *ReceptionistHandler {
	return &ReceptionistHandler{
		DB: db,
	}
}

func (r *ReceptionistHandler) FindPatient(c *gin.Context){

	Request := struct {
		AadharID string `uri:"aadharid"`
	}{}
	err := c.ShouldBindUri(&Request)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	ctx , cancel := context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
	patient , err := repositories.FindPatients(ctx, r.DB , Request.AadharID)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusOK,patient)
}

func (r *ReceptionistHandler) InsertPatient(c *gin.Context){

	var Request struct {
		ID                int       `json:"id,omitempty"`
		Name              string    `json:"name"`
		Phone             string    `json:"phone"`
		Age               int       `json:"age"`
		DOB               string 	`json:"dob"`
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
	}
	err := c.ShouldBindJSON(&Request)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	parsedDob, err := time.Parse("2006-01-02", Request.DOB)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format for DOB"})
        return
    }
	patient := models.Patient{
		Name:              Request.Name,
        Phone:             Request.Phone,
        Age:               Request.Age,
        DOB:               parsedDob, 
        Gender:            Request.Gender,
        EmergencyContact:  Request.EmergencyContact,
        Aadhar:            Request.Aadhar,
        DoctorID:          Request.DoctorID,
        PaymentInfo:       Request.PaymentInfo,
        KnownAllergies:    Request.KnownAllergies,
        Medications:       Request.Medications,
        OtherHealthIssues: Request.OtherHealthIssues,
        DoctorNotes:       Request.DoctorNotes,
        Consent:           Request.Consent,
	}

	ctx , cancel := context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
	patient_id ,err := repositories.InsertPatient(ctx,r.DB,patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,patient_id)
}


func (r *ReceptionistHandler) UpdatePatient(c *gin.Context){
	Request := struct {
		PatientId int `uri:"patientid"`
	}{}
	err := c.ShouldBindUri(&Request)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	var PatientRequest struct {
		ID                int       `json:"id,omitempty"`
		Name              string    `json:"name"`
		Phone             string    `json:"phone"`
		Age               int       `json:"age"`
		DOB               string 	`json:"dob"`
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
	}
	err = c.ShouldBindJSON(&PatientRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if PatientRequest.ID != Request.PatientId {
		c.JSON(http.StatusBadRequest,gin.H{"error":"patient id in body and uri do not match"})
		return
	}
	parsedDob, err := time.Parse("2006-01-02", PatientRequest.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format for DOB"})
		return
	}
	patient := models.Patient{
		ID: 			   PatientRequest.ID,
		Name:              PatientRequest.Name,
        Phone:             PatientRequest.Phone,
        Age:               PatientRequest.Age,
        DOB:               parsedDob, 
        Gender:            PatientRequest.Gender,
        EmergencyContact:  PatientRequest.EmergencyContact,
        Aadhar:            PatientRequest.Aadhar,
        DoctorID:          PatientRequest.DoctorID,
        PaymentInfo:       PatientRequest.PaymentInfo,
        KnownAllergies:    PatientRequest.KnownAllergies,
        Medications:       PatientRequest.Medications,
        OtherHealthIssues: PatientRequest.OtherHealthIssues,
        DoctorNotes:       PatientRequest.DoctorNotes,
        Consent:           PatientRequest.Consent,
	}
	ctx , cancel := context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()

	res ,err := repositories.UpdatePatient(ctx,r.DB,patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,res)
}


func (r *ReceptionistHandler) DeletePatient(c *gin.Context){
	Request := struct {
		PatientId int `uri:"patientid"`
	}{}
	err := c.ShouldBindUri(&Request)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	ctx , cancel := context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
	err = repositories.DeletePatient(ctx,r.DB,Request.PatientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"patient deleted successfully"})
}