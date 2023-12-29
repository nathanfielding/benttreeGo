package handlers

import (
	"benttreeGo/pkg/models"
	"benttreeGo/pkg/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ApartmentHandler struct {
	s *services.DatabaseService
}

func NewApartmentHandler(s *services.DatabaseService) *ApartmentHandler {
	return &ApartmentHandler{s: s}
}

func (h ApartmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe implement to encapsulate all apartment handlers
}

/* -------------------- APARTMENT HANDLER FUNCTIONS -------------------- */
func (h ApartmentHandler) ApartmentList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apartments, err := h.s.FindAllApartments()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(apartments)
	case http.MethodPost:
		var apartment models.Apartment
		if err := json.NewDecoder(r.Body).Decode(&apartment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err := h.s.CreateApartment(apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h ApartmentHandler) ApartmentByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	apartment, err := h.s.FindApartmentByNumber(number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: Implment http.MethodPatch
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(apartment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		if err := json.NewDecoder(r.Body).Decode(apartment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err := h.s.UpdateApartment(apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		err := h.s.DeleteApartment(apartment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
