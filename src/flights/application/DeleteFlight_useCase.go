package application

import "FLIGHTS_API/src/flights/domain"

type DeleteFlight struct {
	repo domain.IFlight
}

func NewDeleteFlight(repo domain.IFlight) *DeleteFlight {
	return &DeleteFlight{repo: repo}
}

func (df *DeleteFlight) Execute(id string) error {
	return df.repo.Delete(id)
}
