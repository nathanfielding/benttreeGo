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

func (r *ApartmentRepository) FindApartmentByNumber(number string) (*models.Apartment, error) {
	var apartment models.Apartment
	query := "SELECT *  FROM apartments WHERE id = ?"
	err := r.db.Get(&apartment, query, number)
	if err != nil {
		return nil, err
	}
	return &apartment, nil
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
