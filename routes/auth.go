package routes

import (
	"net/http"

	"example.com/models"
	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse json"})
		return
	}

	retrievedPassword, err := user.GetUserPasswordByEmail()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid username or password"})
		return
	}

	if !utils.ValidatePassword(user.Password, retrievedPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": utils.SignJwt(user)})

}

func register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot parse json"})
		return
	}

	if user.IsUserExists() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}

	user.Password = utils.HashPassword(user.Password)

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot create user"})
		return
	}

}
