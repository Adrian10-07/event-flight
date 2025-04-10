package domain

type IFlight interface{
	Save(flight *Flight)error
	GetAll()([]Flight,error)
	Delete(id string)error
	Update(id string,flight *Flight)error
	Reservar(userID int, flightID int, seats int) error
}