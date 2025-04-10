package application

import "FLIGHTS_API/src/flights/domain"

type Reserva struct {
	repo domain.IFlight
}

func NewReserva(repo domain.IFlight) *Reserva {
	return &Reserva{repo: repo}
}

func (r *Reserva) Execute(usuarioID int, destinoID int, cantidadBoletos int) error {
	return r.repo.Reservar(usuarioID, destinoID, cantidadBoletos)
}
