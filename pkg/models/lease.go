package models

import (
	"math/big"
	"time"
)

type Lease struct {
	ID            uint      `json:"id" db:"id"`
	TenantID      uint      `json:"tenant_id" db:"tenant_id"`
	TenantName    string    `json:"tenant_name"`
	StartDate     time.Time `json:"start_date" db:"start_date"`
	EndDate       time.Time `json:"end_date" db:"end_date"`
	MonthlyRent   big.Float `json:"monthly_rent" db:"monthly_rent"`
	DepositAmount big.Float `json:"deposit_amount" db:"deposit_amount"`
}

func IsValidLeaseField(field string) bool {
	return (field == "tenant_name" || field == "start_date" || field == "end_date" || field == "monthly_rent" || field == "deposit_amount")
}
