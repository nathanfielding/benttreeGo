package models

import (
	"time"
)

type Lease struct {
	ID              uint      `json:"id" db:"id"`
	TenantID        uint      `json:"tenant_id" db:"tenant_id"`
	TenantName      string    `json:"tenant_name"`
	ApartmentID     uint      `json:"apartment_id" db:"apartment_id"`
	ApartmentNumber string    `json:"apartment_number"`
	StartDate       time.Time `json:"start_date" db:"start_date"`
	EndDate         time.Time `json:"end_date" db:"end_date"`
	MonthlyRent     float32   `json:"monthly_rent" db:"monthly_rent"`
	DepositAmount   float32   `json:"deposit_amount" db:"deposit_amount"`
}

func IsValidLeaseField(field string) bool {
	return (field == "tenant_name" || field == "start_date" || field == "end_date" || field == "monthly_rent" || field == "deposit_amount")
}
