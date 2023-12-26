package models

type Apartment struct {
	ID        uint64
	Number    uint
	Property  string
	Bedrooms  uint
	Occupancy uint
	RentedAs  uint
}

func NewApartment(number uint, property string, bedrooms, occupancy, rentedAs uint) *Apartment {
	return &Apartment{
		Number:    number,
		Property:  property,
		Bedrooms:  bedrooms,
		Occupancy: occupancy,
		RentedAs:  rentedAs,
	}
}
