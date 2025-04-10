package controllers

import (
	"FLIGHTS_API/src/flights/application"
	"FLIGHTS_API/src/flights/infraestructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllFlightsHandler(c *gin.Context) {
	repo := infraestructure.NewMySQLFlightRepository()
	useCase := application.NewGetAllFlights(repo)

	flights, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los vuelos"})
		return
	}

	c.JSON(http.StatusOK, flights) 
}
