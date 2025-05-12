package main

import (
	"echo/db"
	"echo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()
	routes.RegisterRoutes(r)
	routes.LoginRoute(r)
	r.Run(":8080")
}
