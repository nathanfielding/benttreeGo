package models

type Apartment struct {
	ID        uint64 `json:"id" db:"id"`
	Number    string `json:"number" db:"number"`
	Property  string `json:"property" db:"property"`
	Bedrooms  uint   `json:"bedrooms" db:"bedrooms"`
	Occupancy uint   `json:"occupancy" db:"occupancy"`
	RentedAs  uint   `json:"rented_as" db:"rented_as"`
}
