package main

import (
	"echo/db"
	"echo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
