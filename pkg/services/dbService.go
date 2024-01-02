package services

import (
	"benttreeGo/pkg/models"

	"github.com/jmoiron/sqlx"
)

type DatabaseService struct {
	db *sqlx.DB
}

func NewDatabaseService(db *sqlx.DB) *DatabaseService {
	return &DatabaseService{db: db}
}

/* -------------------- APARTMENT HELPER FUNCTIONS -------------------- */
func (s *DatabaseService) FindAllApartments() ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM Apartments"
	err := s.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

func (s *DatabaseService) FindApartmentByNumber(number string) (*models.Apartment, error) {
	var apartment models.Apartment
	query := "SELECT * FROM Apartments WHERE number = $1"
	err := s.db.Get(&apartment, query, number)
	if err != nil {
		return nil, err
	}
	return &apartment, nil
}

func (s *DatabaseService) FindApartmentIDByNumber(number string) (uint64, error) {
	var id uint64
	query := "SELECT id from Apartments WHERE number = $1"
	if err := s.db.Get(&id, query, number); err != nil {
		return 0, err
	}
	return id, nil

}

func (s *DatabaseService) FindApartmentByBedrooms(bedrooms uint) ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM Apartments WHERE bedrooms = $1"
	err := s.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

func (s *DatabaseService) CreateApartment(a models.Apartment) error {
	query := "INSERT INTO Apartments (number, property, bedrooms, occupancy, rented_as) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatabaseService) UpdateApartment(a *models.Apartment) error {
	query := "UPDATE Apartments SET number = $1, property = $2, bedrooms = $3, occupancy = $4, rented_as = $5 WHERE number = $1"
	_, err := s.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatabaseService) DeleteApartment(a *models.Apartment) error {
	query := "DELETE FROM Apartments WHERE number = $1"
	_, err := s.db.Exec(query, a.Number)
	if err != nil {
		return err
	}
	return nil
}

/* -------------------- TENANT HELPER FUNCTIONS -------------------- */
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
	query := "SELECT * FROM Tenants JOIN Apartments ON Tenants.apartment_id = Apartments.id WHERE Apartments.number = $1"
	err := s.db.Select(tenants, query, number)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

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

func (s DatabaseService) DeleteTenant(t *models.Tenant) error {
	query := "DELETE FROM Tenants WHERE name = $1"
	_, err := s.db.Exec(query, t.Name)
	if err != nil {
		return err
	}
	return nil
}
