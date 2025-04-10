package controllers

import (
	"FLIGHTS_API/src/flights/application"
	"FLIGHTS_API/src/flights/domain"
	"FLIGHTS_API/src/flights/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateFlightHandler(c *gin.Context) {
	id := c.Param("id") 

	var flight domain.Flight

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	repo := infraestructure.NewMySQLFlightRepository()
	useCase := application.NewUpdateFlight(repo)

	if err := useCase.Execute(id, flight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el vuelo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vuelo actualizado con éxito"})
}
