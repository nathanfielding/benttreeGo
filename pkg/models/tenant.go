package models

// part of post/patch/put payload:
// apartment_number", ?(lease_id), name, email, phone_number, home_address, is_renewing
type Tenant struct {
	ID              uint   `json:"id" db:"id"`
	ApartmentID     uint   `json:"apartment_id" db:"apartment_id"`
	ApartmentNumber string `json:"apartment_number"`
	LeaseID         *uint  `json:"lease_id" db:"lease_id"` // nullable
	Name            string `json:"name" db:"name"`
	Email           string `json:"email" db:"email"`
	PhoneNumber     string `json:"phone_number" db:"phone_number"`
	HomeAddress     string `json:"home_address" db:"home_address"`
	IsRenewing      bool   `json:"is_renewing" db:"is_renewing"`
}

func IsValidTenantField(field string) bool {
	return (field == "apartment_number" || field == "name" || field == "email" || field == "phone_number" || field == "home_address" || field == "is_renewing")
}
