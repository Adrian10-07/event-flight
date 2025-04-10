package routes

import (
	"FLIGHTS_API/src/flights/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	flights := router.Group("/flights")
	{
		flights.POST("/", controllers.CreateFlightHandler)
		flights.GET("/", controllers.GetAllFlightsHandler)  
		flights.PUT("/:id", controllers.UpdateFlightHandler) 
		flights.DELETE("/:id", controllers.DeleteFlightHandler)
 	}
}
