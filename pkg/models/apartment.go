package models

type Apartment struct {
	ID        uint   `json:"id" db:"id"`
	Number    string `json:"number" db:"number"`
	Property  string `json:"property" db:"property"`
	Bedrooms  uint   `json:"bedrooms" db:"bedrooms"`
	Occupancy uint   `json:"occupancy" db:"occupancy"`
	RentedAs  uint   `json:"rented_as" db:"rented_as"`
}

func ValidApartmentField(field string) bool {
	switch field {
	case "number":
	case "property":
	case "bedrooms":
	case "occupancy":
	case "rented_as":
		return true
	}
	return false
}
