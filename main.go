package main

import (
	"net/http"

	"github.com/akeemkabiru/booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvent(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(c *gin.Context) {
	var events models.Event

	err := c.ShouldBindJSON(events)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "could not parse data"})
		return
	}

	events.ID = 1
	events.UserId = 1
	events.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "data successfully parsed"})
}
