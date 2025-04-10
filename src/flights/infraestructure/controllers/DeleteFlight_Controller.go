package controllers

import (
	"net/http"
	"FLIGHTS_API/src/flights/application"
	"FLIGHTS_API/src/flights/infraestructure"
	"github.com/gin-gonic/gin"

)

func DeleteFlightHandler(c *gin.Context) {
	id := c.Param("id")

	repo := infraestructure.NewMySQLFlightRepository()
	useCase := application.NewDeleteFlight(repo)

	if err := useCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar vuelo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vuelo eliminado"})
}

