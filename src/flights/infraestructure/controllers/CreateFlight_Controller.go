package controllers

import (
	"FLIGHTS_API/src/flights/application"
	"FLIGHTS_API/src/flights/domain"
	"FLIGHTS_API/src/flights/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func CreateFlightHandler(c *gin.Context) {
	var flight domain.Flight

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	repo := infraestructure.NewMySQLFlightRepository()
	eventPublisher, err := infraestructure.NewRabbitMQPublisher()
	if err != nil {
		log.Printf("Error conectando con RabbitMQ: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en el servicio de mensajería"})
		return
	}

	useCase := application.NewCreateFlight(repo, eventPublisher)

	err = useCase.Execute(flight)
	if err != nil {
		log.Printf("Error al crear el vuelo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error desconocido"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Vuelo creado con éxito"})
}
