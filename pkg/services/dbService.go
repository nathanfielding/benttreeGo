package services

import (
	"benttreeGo/pkg/models"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type DatabaseService struct {
	db *sqlx.DB
}

func NewDatabaseService(db *sqlx.DB) DatabaseService {
	return DatabaseService{db: db}
}

/* -------------------- APARTMENT HELPER FUNCTIONS -------------------- */
func (s DatabaseService) CreateApartment(a models.Apartment) error {
	query := "INSERT INTO Apartments (number, property, bedrooms, occupancy, rented_as) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) PutApartment(a *models.Apartment) error {
	query := "UPDATE Apartments SET number = $1, property = $2, bedrooms = $3, occupancy = $4, rented_as = $5 WHERE number = $1"
	_, err := s.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) PatchApartment(number, field string, rawValue interface{}) error {
	if !models.IsValidApartmentField(field) {
		return fmt.Errorf("invalid updatable apartment field name")
	}
	query := fmt.Sprintf("UPDATE Apartments SET %s = $1 WHERE number = $2", field)
	_, err := s.db.Exec(query, rawValue, number)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) DeleteApartment(a *models.Apartment) error {
	query := "DELETE FROM Apartments WHERE number = $1"
	_, err := s.db.Exec(query, a.Number)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) FindAllApartments() ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM Apartments"
	err := s.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

func (s DatabaseService) FindApartmentByNumber(number string) (*models.Apartment, error) {
	var apartment models.Apartment
	query := "SELECT * FROM Apartments WHERE number = $1"
	err := s.db.Get(&apartment, query, number)
	if err != nil {
		return nil, err
	}
	return &apartment, nil
}

func (s DatabaseService) FindApartmentIDByNumber(number string) (uint, error) {
	var id uint
	query := "SELECT id from Apartments WHERE number = $1"
	if err := s.db.Get(&id, query, number); err != nil {
		return 0, err
	}
	return id, nil

}

func (s DatabaseService) FindApartmentByBedrooms(bedrooms uint) ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM Apartments WHERE bedrooms = $1"
	err := s.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

/* -------------------- TENANT HELPER FUNCTIONS -------------------- */
func (s DatabaseService) CreateTenant(t *models.Tenant) error {
	apartmentID, err := s.FindApartmentIDByNumber(t.ApartmentNumber) // may have to do same with leaseID
	if err != nil {
		return err
	}
	query := `INSERT INTO Tenants (apartment_id, lease_id, name, email, phone_number, home_address, is_renewing)
				VALUES ($1, $2, $3, $4, $5, $6, $7)` // eventually maybe add lease_id
	_, err = s.db.Exec(query, apartmentID, nil, t.Name, t.Email, t.PhoneNumber, t.HomeAddress, t.IsRenewing)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) PutTenant(t models.Tenant) error {
	apartmentID, err := s.FindApartmentIDByNumber(t.ApartmentNumber)
	if err != nil {
		return err
	}
	query := `UPDATE Tenants SET apartment_id = $1, name = $2, email = $3, phone_number = $4,
				home_address = $5, is_renewing = $6 WHERE name = $2`
	_, err = s.db.Exec(query, apartmentID, t.Name, t.Email, t.PhoneNumber, t.HomeAddress, t.IsRenewing)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) PatchTenant(name, field string, rawValue interface{}) error {
	if !models.IsValidTenantField(field) {
		return fmt.Errorf("invalid updatable tenant field name")
	}
	// case where field is ApartmentNumber (not stored in db)
	if field == "apartment_number" {
		apartmentID, err := s.FindApartmentIDByNumber(rawValue.(string))
		if err != nil {
			return err
		}
		rawValue = apartmentID
	}

	query := fmt.Sprintf("UPDATE Tenants SET %s = $1 WHERE name = $2", field)
	_, err := s.db.Exec(query, rawValue, name)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) DeleteTenant(t *models.Tenant) error {
	query := "DELETE FROM Tenants WHERE name = $1"
	_, err := s.db.Exec(query, t.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) FindAllTenants() ([]models.Tenant, error) {
	var tenants []models.Tenant
	query := `SELECT Tenants.*, Apartments.number AS ApartmentNumber FROM Tenants
				JOIN Apartments ON Tenants.apartment_id = Apartments.id`
	err := s.db.Select(&tenants, query)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

func (s DatabaseService) FindTenantByName(name string) (*models.Tenant, error) {
	var tenant models.Tenant
	query := `SELECT Tenants.*, Apartments.number AS ApartmentNumber FROM Tenants
				JOIN Apartments ON Tenants.apartment_id = Apartments.id
				WHERE name = $1`
	err := s.db.Get(&tenant, query, name)
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s DatabaseService) FindTenantsByApartmentNumber(number string) ([]models.Tenant, error) {
	var tenants []models.Tenant
	query := `SELECT Tenants.*, Apartments.number AS ApartmentNumber FROM Tenants 
				JOIN Apartments ON Tenants.apartment_id = Apartments.id WHERE Apartments.number = $1`
	err := s.db.Select(tenants, query, number)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

func (s DatabaseService) FindTenantIDByName(name string) (uint, error) {
	var id uint
	query := "SELECT id FROM Tenants WHERE name = $1"
	if err := s.db.Get(&id, query, name); err != nil {
		return 0, err
	}
	return id, nil
}

/* -------------------- LEASE HELPER FUNCTIONS -------------------- */
func (s DatabaseService) CreateLease(l models.Lease) error {
	tenantID, err := s.FindTenantIDByName(l.TenantName)
	if err != nil {
		return nil
	}
	startDate := l.StartDate.Format("2000-01-01")
	endDate := l.EndDate.Format("2000-01-01")

	query := `INSERT INTO Leases (tenant_id, start_date, end_date, monthly_rent, deposit_amount)
				VALUES ($1, $2, $3, $4, $5)`

	_, err = s.db.Exec(query, tenantID, startDate, endDate, l.MonthlyRent, l.DepositAmount)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) PutLease(l models.Lease) error {
	tenantID, err := s.FindTenantIDByName(l.TenantName)
	if err != nil {
		return err
	}
	query := `"UPDATE Leases SET tenant_id = $1, start_date = $2, end_date = $3,
				monthly_rent = $4, deposit_amount = $5`
	_, err = s.db.Exec(query, tenantID, l.StartDate, l.EndDate, l.MonthlyRent, l.DepositAmount)
	if err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) PatchLease(name, field string, rawValue interface{}) error {
	if !models.IsValidLeaseField(field) {
		return fmt.Errorf("invalid updatable lease field name")
	}
	tenantID, err := s.FindTenantIDByName(name)
	if err != nil {
		return err
	}
	// case where field is TenantName (not stored in db)
	if field == "tenant_name" {
		rawValue = tenantID
	}

	query := fmt.Sprintf("UPDATE Leases SET %s = $1 WHERE tenant_id = $2", field)
	if _, err := s.db.Exec(query, rawValue, tenantID); err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) DeleteLease(l models.Lease) error {
	query := "DELETE FROM Leases WHERE tenant_id = $1"
	if _, err := s.db.Exec(query, l.TenantID); err != nil {
		return err
	}
	return nil
}

func (s DatabaseService) FindAllLeases() ([]models.Lease, error) {
	var leases []models.Lease
	query := `SELECT Leases.*, Tenants.name AS TenantName FROM Leases
				JOIN Tenants ON Leases.tenant_id = Tenants.id`
	if err := s.db.Select(&leases, query); err != nil {
		return nil, err
	}
	return leases, nil
}

func (s DatabaseService) FindLeaseByName(name string) (models.Lease, error) {
	var lease models.Lease
	tenantID, err := s.FindTenantIDByName(name)
	if err != nil {
		return models.Lease{}, err
	}
	query := `SELECT Leases.*, Tenants.name AS TenantName FROM Leases
				JOIN Tenants ON Leases.tenant_id = Tenants.id WHERE Leases.tenant_id = $1`
	err = s.db.Get(&lease, query, tenantID)
	if err != nil {
		return models.Lease{}, err
	}
	return lease, nil
}

/* -------------------- MISCELLANEOUS HELPER FUNCTIONS -------------------- */
func (s DatabaseService) FindApartmentsByEndDate(endDate time.Time) ([]models.Apartment, error) {
	var apartmentIDs []uint
	query := `SELECT Tenant.apartment_id FROM Tenant JOIN Leases
				ON Leases.tenant_id = Tenant.id WHERE Leases.end_date <= $1
				AND Tenant.is_renwing = FALSE`
	err := s.db.Select(apartmentIDs, query, endDate)
	if err != nil {
		return nil, err
	}
	var apartments []models.Apartment
	query = "SELECT * FROM Apartments WHERE id = ANY ($1)"
	err = s.db.Select(apartments, query, apartmentIDs)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}
