package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"eventgo.com/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get events."})
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect request body."})
	}

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save."})
	}
	c.JSON(http.StatusOK, event)
}

func getSingleEvent(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect request params."})
	}

	event, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event with that id."})
	}

	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect request params."})
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event with that id."})
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect request body."})
	}

	updatedEvent.ID = eventId

	fmt.Println("EVENT: ", updatedEvent)
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event with that id."})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Successful"})
}
