package application

import "FLIGHTS_API/src/flights/domain"

type UpdateFlight struct {
	repo domain.IFlight
}

func NewUpdateFlight(repo domain.IFlight) *UpdateFlight {
	return &UpdateFlight{repo: repo}
}

func (uf *UpdateFlight) Execute(id string, flight domain.Flight) error {
	return uf.repo.Update(id, &flight)
}
