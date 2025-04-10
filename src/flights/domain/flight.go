package domain

type Flight struct {
	ID             int    
	Origin         string 
	Destination    string 
	TotalSeats     int    
	AvailableSeats int    
	Status         string 
}
