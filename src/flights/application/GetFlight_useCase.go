package application

import "FLIGHTS_API/src/flights/domain"

type GetAllFlights struct {
	repo domain.IFlight
}

func NewGetAllFlights(repo domain.IFlight) *GetAllFlights {
	return &GetAllFlights{repo: repo}
}

func (gf *GetAllFlights) Execute() ([]domain.Flight, error) {
	return gf.repo.GetAll()
}
