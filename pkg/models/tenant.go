package models

// https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267

type Tenant struct {
	ID              uint64  `json:"id" db:"id"`
	ApartmentID     uint64  `json:"apartment_id" db:"apartment_id"`
	ApartmentNumber string  `json:"apartment_number"`       // not stored in database
	LeaseID         *uint64 `json:"lease_id" db:"lease_id"` // allows for null
	Name            string  `json:"name" db:"name"`
	Email           string  `json:"email" db:"email"`
	PhoneNumber     string  `json:"phone_number" db:"phone_number"`
	HomeAddress     string  `json:"home_address" db:"home_address"`
	IsRenewing      bool    `json:"is_renewing" db:"is_renewing"`
}
