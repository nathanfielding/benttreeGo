package models

type Apartment struct {
	ID        uint   `json:"id" db:"id"`
	Number    string `json:"number" db:"number"`
	Property  string `json:"property" db:"property"`
	Bedrooms  uint   `json:"bedrooms" db:"bedrooms"`
	Occupancy uint   `json:"occupancy" db:"occupancy"`
	RentedAs  uint   `json:"rented_as" db:"rented_as"`
}

func IsValidApartmentField(field string) bool {
	return (field == "id" || field == "number" || field == "property" || field == "bedrooms" || field == "occupancy" || field == "rented_as")
}
