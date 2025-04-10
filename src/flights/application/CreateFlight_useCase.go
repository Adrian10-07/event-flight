package application

import "FLIGHTS_API/src/flights/domain"
import
"fmt"


type CreateFlight struct {
	repo           domain.IFlight
	eventPublisher domain.EventPublisher
}

func NewCreateFlight(repo domain.IFlight, eventPublisher domain.EventPublisher) *CreateFlight {
	return &CreateFlight{
		repo:           repo,
		eventPublisher: eventPublisher,
	}
}

func (cf *CreateFlight) Execute(f domain.Flight) error {
	err := cf.repo.Save(&f)
	if err != nil {
		return fmt.Errorf("error guardando el vuelo: %v", err)
	}

	err = cf.eventPublisher.Publish("flight_created", f)
	if err != nil {
		return fmt.Errorf("error publicando el evento: %v", err)
	}

	return nil
}
