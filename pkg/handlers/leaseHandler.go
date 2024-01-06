package handlers

import (
	"benttreeGo/pkg/models"
	"benttreeGo/pkg/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type LeaseHandler struct {
	s services.DatabaseService
}

func NewLeaseHandler(s services.DatabaseService) LeaseHandler {
	return LeaseHandler{s: s}
}

func (h LeaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe implement to encapsulate all lease handlers
}

/* -------------------- LEASE HANDLER FUNCTIONS -------------------- */
func (h LeaseHandler) LeaseList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		leases, err := h.s.FindAllLeases()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = json.NewEncoder(w).Encode(leases); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var lease models.Lease
		if err := json.NewDecoder(r.Body).Decode(&lease); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := h.s.CreateLease(lease); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h LeaseHandler) LeaseByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lease, err := h.s.FindLeaseByName(vars["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(lease); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		if err := json.NewDecoder(r.Body).Decode(&lease); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := h.s.PutLease(lease); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPatch:
		updates := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for k, v := range updates {
			if err := h.s.PatchLease(vars["name"], k, v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	case http.MethodDelete:
		if err := h.s.DeleteLease(lease); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
