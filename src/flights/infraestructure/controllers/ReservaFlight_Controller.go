package controllers

/* import (
	"fmt"
	"net/http"
	"FLIGHTS_API/src/flights/infraestructure"
	"FLIGHTS_API/src/flights/infraestructure/adaptadores"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Estructura para recibir la solicitud de reserva
type ReservaRequest struct {
	CantidadBoletos int `json:"cantidad_boletos"`
}

func RealizarReserva(c *gin.Context) {
	// Obtener parámetros desde la URL
	usuarioIDParam := c.Param("usuarioId")
	usuarioID, err := strconv.Atoi(usuarioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuarioId debe ser un número"})
		return
	}

	vueloIDParam := c.Param("vueloId")
	vueloID, err := strconv.Atoi(vueloIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vueloId debe ser un número"})
		return
	}

	// Obtener datos del body (cantidad de boletos)
	var req ReservaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar JSON"})
		return
	}

	// Crear la reserva
	repo := infraestructure.NewMySQLFlightRepository()

	// Lógica de reserva (verificar disponibilidad y actualizar la base de datos)
	if err := repo.Reservar(usuarioID, vueloID, req.CantidadBoletos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No se pudo procesar la reserva: %v", err)})
		return
	}

	// Publicar la reserva en RabbitMQ (si se usa para otros procesos asíncronos)
// Crear la reserva
reserva := adaptadores.Reservation{
	ID:        int(usuarioID + vueloID + int(time.Now().UnixNano()%1000)), // Generar un ID único numérico
	UserID:    usuarioID,
	FlightID:  vueloID,
	Seats:     req.CantidadBoletos,
	Status:    "pending", // Por defecto, el estado es 'pending'
	Timestamp: time.Now().Format(time.RFC3339),
}


	success, err := adaptadores.PublishReservation(reserva)
	if err != nil || !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la publicación de la reserva"})
		return
	}

	// Respuesta de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Reserva realizada con éxito"})
} */
