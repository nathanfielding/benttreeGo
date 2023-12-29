package handlers

import (
	"benttreeGo/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type ApartmentHandler struct {
	db *sqlx.DB
}

func NewApartmentHandler(db *sqlx.DB) *ApartmentHandler {
	return &ApartmentHandler{db: db}
}

func (c ApartmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe implement to encapsulate all apartment handlers
}

/* -------------------- APARTMENT HANDLER FUNCTIONS -------------------- */
func (c ApartmentHandler) ApartmentList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apartments, err := c.FindAllApartments()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(apartments)
	case http.MethodPost:
		var apartment models.Apartment
		json.NewDecoder(r.Body).Decode(&apartment)
		err := c.CreateApartment(&apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c ApartmentHandler) ApartmentByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	apartment, err := c.FindApartmentByNumber(number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// TODO: Implment http.MethodPatch
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(apartment)
	case http.MethodPut:
		json.NewDecoder(r.Body).Decode(apartment)
		err := c.UpdateApartment(apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		err := c.DeleteApartment(apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

/* -------------------- APARTMENT HELPER FUNCTIONS -------------------- */
func (c *ApartmentHandler) FindAllApartments() ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM Apartments"
	err := c.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}

func (c *ApartmentHandler) FindApartmentByNumber(number string) (*models.Apartment, error) {
	var apartment models.Apartment
	query := "SELECT * FROM Apartments WHERE number = $1"
	err := c.db.Get(&apartment, query, number)
	if err != nil {
		return nil, err
	}
	return &apartment, nil
}

func (c *ApartmentHandler) CreateApartment(a *models.Apartment) error {
	query := "INSERT INTO Apartments (number, property, bedrooms, occupancy, rented_as) VALUES ($1, $2, $3, $4, $5)"
	_, err := c.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (c *ApartmentHandler) UpdateApartment(a *models.Apartment) error {
	query := "UPDATE Apartments SET number = $1, property = $2, bedrooms = $3, occupancy = $4, rented_as = $5 WHERE number = $1"
	_, err := c.db.Exec(query, a.Number, a.Property, a.Bedrooms, a.Occupancy, a.RentedAs)
	if err != nil {
		return err
	}
	return nil
}

func (c *ApartmentHandler) DeleteApartment(a *models.Apartment) error {
	query := "DELETE FROM Apartments WHERE number = $1"
	_, err := c.db.Exec(query, a.Number)
	if err != nil {
		return err
	}
	return nil
}

func (c *ApartmentHandler) FindAparmentByBedrooms(bedrooms uint) ([]models.Apartment, error) {
	var apartments []models.Apartment
	query := "SELECT * FROM Apartments WHERE bedrooms = $1"
	err := c.db.Select(&apartments, query)
	if err != nil {
		return nil, err
	}
	return apartments, nil
}
