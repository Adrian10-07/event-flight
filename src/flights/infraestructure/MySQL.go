package infraestructure

import (
	"FLIGHTS_API/src/flights/domain"
	"FLIGHTS_API/src/core"
	"fmt"
	"log"
	"strconv"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLFlightRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(f *domain.Flight) error {
	if f.TotalSeats <= 0 {
		return fmt.Errorf("total_seats debe ser mayor que 0")
	}
	if f.AvailableSeats < 0 {
		return fmt.Errorf("available_seats no puede ser negativo")
	}
	if f.Status != "scheduled" && f.Status != "full" && f.Status != "cancelled" {
		return fmt.Errorf("status debe ser uno de 'scheduled', 'full', 'cancelled'")
	}

	query := "INSERT INTO flights (origin, destination, total_seats, available_seats, status) VALUES (?, ?, ?, ?, ?)"
	_, err := r.conn.DB.Exec(query, f.Origin, f.Destination, f.TotalSeats, f.AvailableSeats, f.Status)
	if err != nil {
		log.Printf("Error al guardar vuelo: %v", err)
		return err
	}

	log.Println("Vuelo guardado exitosamente")
	return nil
}

func (r *MySQLRepository) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	query := "DELETE FROM flights WHERE id = ?"
	_, err = r.conn.DB.Exec(query, intID)
	return err
}

func (r *MySQLRepository) Update(id string, f *domain.Flight) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	query := "UPDATE flights SET origin = ?, destination = ?, total_seats = ?, available_seats = ?, status = ? WHERE id = ?"
	_, err = r.conn.DB.Exec(query, f.Origin, f.Destination, f.TotalSeats, f.AvailableSeats, f.Status, intID)
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Flight, error) {
	query := "SELECT id, origin, destination, total_seats, available_seats, status FROM flights"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flights []domain.Flight
	for rows.Next() {
		var flight domain.Flight
		if err := rows.Scan(&flight.ID, &flight.Origin, &flight.Destination, &flight.TotalSeats, &flight.AvailableSeats, &flight.Status); err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	return flights, nil
}

func (r *MySQLRepository) GetByID(id string) (*domain.Flight, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	query := "SELECT id, origin, destination, total_seats, available_seats, status FROM flights WHERE id = ?"
	row := r.conn.DB.QueryRow(query, intID)

	var flight domain.Flight
	err = row.Scan(&flight.ID, &flight.Origin, &flight.Destination, &flight.TotalSeats, &flight.AvailableSeats, &flight.Status)
	if err != nil {
		return nil, err
	}
	return &flight, nil
}

func (r *MySQLRepository) Reservar(userID int, flightID int, seats int) error {
	var availableSeats int
	query := "SELECT available_seats FROM flights WHERE id = ?"
	err := r.conn.DB.QueryRow(query, flightID).Scan(&availableSeats)
	if err != nil {
		return fmt.Errorf("error al obtener disponibilidad del vuelo: %v", err)
	}

	if availableSeats < seats {
		return fmt.Errorf("no hay suficientes asientos disponibles")
	}

	insertQuery := "INSERT INTO reservations (user_id, flight_id, seats, status) VALUES (?, ?, ?, 'pending')"
	_, err = r.conn.DB.Exec(insertQuery, userID, flightID, seats)
	if err != nil {
		return fmt.Errorf("error al registrar la reserva: %v", err)
	}

	updateQuery := "UPDATE flights SET available_seats = available_seats - ? WHERE id = ?"
	_, err = r.conn.DB.Exec(updateQuery, seats, flightID)
	if err != nil {
		return fmt.Errorf("error al actualizar asientos disponibles: %v", err)
	}

	log.Println("Reserva realizada con éxito")
	return nil
}