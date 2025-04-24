package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/Somvaded/assessment/repositories"
	"github.com/Somvaded/assessment/utils"
	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	DB *sql.DB
}


func NewUserHandler (db *sql.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func(h *UserHandler) Login( ctx *gin.Context){
	var Request struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBindJSON(&Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return 
	}
	c, cancel := context.WithTimeout(ctx.Request.Context(), 10* time.Second)
	defer cancel()
	user , doctor, receptionist, err := repositories.FindUserByEmail(c,h.DB,Request.Email,Request.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return 
	}
	
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	if user.Role == "doctor"{
		ctx.SetCookie("auth_token",token,3600*24,"/","",false,true)
		ctx.JSON(http.StatusOK,doctor)
	} else{
		ctx.SetCookie("auth_token",token,3600*24,"/","",false,true)
		ctx.JSON(http.StatusOK,receptionist)
	}
}