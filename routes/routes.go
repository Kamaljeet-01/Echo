package routes

import (
	model "echo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", func(c *gin.Context) {
		var req struct {
			Username string `json : "username"`
			Password string `json: "password"`
		}
		// Parse and bind the incoming JSON request to req struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		// Check if the username already exists in the database
		exist, err := model.Checkuserexist(req.Username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if exist {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		//if username does not exists
		err = model.InsertUser(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	//LOGIN ROUTE

}
func LoginRoute(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json : "username"`
			Password string `json:"passowrd`
		}
		//Bind login json
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		//Call function to verify credentials
		valid, err := model.CheckUserCred(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//if password doesnt match
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successfull"})

	})
}
