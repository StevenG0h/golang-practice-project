package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(c *gin.Context) {
	param := c.Param("eventId")

	eventId, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt64("userId")

	registerEvent := models.Register{
		UserId:  int64(userId),
		EventId: eventId,
	}

	if registerEvent.IsRegistrationExists() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "registration already exists"})
		return
	}

	err = registerEvent.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot register to event"})
		return
	}
}

func cancelRegisterEvent(c *gin.Context) {
	param := c.Param("eventId")

	eventId, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userId := c.GetInt64("userId")

	registerEvent := models.Register{
		UserId:  int64(userId),
		EventId: eventId,
	}

	if registerEvent.IsRegistrationExists() != true {
		c.JSON(http.StatusBadRequest, gin.H{"message": "registration not exists"})
		return
	}

	err = registerEvent.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to remove registration"})
		return
	}
}
