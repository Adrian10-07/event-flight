package main

import (
	"FLIGHTS_API/src/flights/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080") 
}
