package routes

import (
	"database/sql"

	"github.com/Somvaded/assessment/handlers"
	"github.com/Somvaded/assessment/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB){
	userHandlers := handlers.NewUserHandler(db)
	receptionistHandlers := handlers.NewReceptionistHandler(db)
	doctorHandlers := handlers.NewDoctorHandler(db)
	
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
	
	// User login 
	userPath := router.Group("/api")
	userPath.POST("/login",userHandlers.Login)

	//recetionist routes
	receptionistPath := router.Group("/api/receptionist",middlewares.Protect(),middlewares.CheckRole("receptionist"))
	receptionistPath.POST("/",receptionistHandlers.InsertPatient)
	receptionistPath.GET("/:aadharid",receptionistHandlers.FindPatient)
	receptionistPath.PUT("/:patientid",receptionistHandlers.UpdatePatient)
	receptionistPath.DELETE("/:patientid",receptionistHandlers.DeletePatient)

	//doctor routes
	doctorPath := router.Group("/api/doctor",middlewares.Protect(),middlewares.CheckRole("doctor"))
	doctorPath.GET("/myPatients",doctorHandlers.GetAllPatientsAssigned)
	doctorPath.PATCH("/:patientid",doctorHandlers.UpdatePatientDetail)
} 