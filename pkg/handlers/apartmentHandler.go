package handlers

import (
	"benttreeGo/pkg/models"
	"benttreeGo/pkg/repositories"
	"encoding/json"
	"net/http"
)

type ApartmentHandler struct {
	aptRepo *repositories.ApartmentRepository
}

func NewApartmentHandler(aptRepo *repositories.ApartmentRepository) *ApartmentHandler {
	return &ApartmentHandler{aptRepo: aptRepo}
}

func (c ApartmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (c ApartmentHandler) ApartmentList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apartments, err := c.aptRepo.FindAllApartments()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(apartments)
	case http.MethodPost:
		var apartment models.Apartment
		json.NewDecoder(r.Body).Decode(&apartment)
		err := c.aptRepo.CreateApartment(&apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c ApartmentHandler) ApartmentByNumber(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement and determine router for custom URL paths

}
