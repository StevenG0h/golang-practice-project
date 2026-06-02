package routes

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse json"})
		return
	}

	if user.IsUserExists() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create user"})
		return
	}
}
