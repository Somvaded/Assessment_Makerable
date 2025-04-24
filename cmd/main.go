package main

import (
	"fmt"

	"github.com/Somvaded/assessment/config"
	"github.com/Somvaded/assessment/db"
	"github.com/Somvaded/assessment/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default();
	db := db.ConnectDatabase(conn.DBUrl);
	routes.RegisterRoutes(r,db)

	if conn.Port == "" {
		conn.Port = "8080"
	}
	r.Run(fmt.Sprintf(":%s",conn.Port))
}