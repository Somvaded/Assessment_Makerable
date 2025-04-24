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

type DoctorHandler struct {
	DB *sql.DB
}

func NewDoctorHandler(db *sql.DB) *DoctorHandler {
	return &DoctorHandler{
		DB: db,
	}
}

func (d *DoctorHandler) GetAllPatientsAssigned(c *gin.Context) {
	doctor_id ,exists := c.Get("user_id")
	if doctor_id == nil || !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	ctx , cancel := context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()

	patients, err := repositories.FindPatientsByDoctorID(ctx, d.DB, doctor_id.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (d *DoctorHandler) UpdatePatientDetail(c *gin.Context)  {
	var Request struct{
		PatientId int `uri:"patientid"`
	}

	err := c.ShouldBindUri(&Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid patient ID",
		})
		return
	}

	var updateData models.DocPatientUpdate

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	ctx , cancel := context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
	updatedPatient,err := repositories.UpdateMedicalInfo(ctx, d.DB, Request.PatientId, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,updatedPatient)
}