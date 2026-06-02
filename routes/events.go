package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, events)
}

func GetEventById(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	event, err := models.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, event)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	event.UserID = c.GetInt("userId")

	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	event.Save()

	c.JSON(http.StatusCreated, event)
}

func UpdateEvent(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = models.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	event.ID = int(id)

	err = event.UpdateById()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, event)
}

func DeleteEvent(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = models.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	err = models.DeleteById(int(id))

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event deleted",
	})
}
