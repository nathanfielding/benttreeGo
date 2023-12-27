package repositories

import (
	"benttreeGo/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ApartmentRepository struct {
	db *sqlx.DB
}

func NewApartmentRepository(db *sqlx.DB) *ApartmentRepository {
	return &ApartmentRepository{db: db}
}

func (r *ApartmentRepository) FindAllApartments() ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM apartments"
	err := r.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

func (r *ApartmentRepository) FindApartmentByNumber(number string) (*models.Apartment, error) {
	var apartment models.Apartment
	query := "SELECT *  FROM apartments WHERE id = ?"
	err := r.db.Get(&apartment, query, number)
	if err != nil {
		return nil, err
	}
	return &apartment, nil
}

func (r *ApartmentRepository) CreateApartment(a *models.Apartment) error {
	query := "INSERT INTO apartments (number, property, bedrooms, occupancy, rented_as) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (r *ApartmentRepository) UpdateApartment(a *models.Apartment) error {
	query := "UPDATE apartments SET number = ?, property = ?, bedrooms = ?, occupancy = ?, rented_as = ? WHERE number = ?"
	_, err := r.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs, a.Number)
	if err != nil {
		return err
	}
	return nil
}

func (r *ApartmentRepository) DeleteApartment(a *models.Apartment) error {
	query := "DELETE FROM apartments WHERE number = ?"
	err := r.db.QueryRow(query, a.Number).Scan()
	if err != nil {
		return err
	}
	return nil
}

func (r *ApartmentRepository) FindAparmentByBedrooms(bedrooms uint) ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM apartments WHERE bedrooms = ?"
	err := r.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}
