package main

import (
	"echo/db"
	model "echo/models"
	"echo/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()
	routes.RegisterRoutes(r)
	routes.LoginRoute(r)
	r.DELETE("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		ok, err := model.Deleteuser(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})
	r.Run(":8080")
}
