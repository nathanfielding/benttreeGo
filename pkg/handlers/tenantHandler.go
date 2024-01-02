package handlers

import (
	"benttreeGo/pkg/models"
	"benttreeGo/pkg/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TenantHandler struct {
	s *services.DatabaseService
}

func NewTenantHandler(s *services.DatabaseService) *TenantHandler {
	return &TenantHandler{s: s}
}

func (h TenantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe implement to encapsulate all tenant handlers
}

/* -------------------- TENANT HANDLER FUNCTIONS -------------------- */
func (h TenantHandler) TenantList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tenants, err := h.s.FindAllTenants()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err = json.NewEncoder(w).Encode(tenants); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var tenant models.Tenant
		if err := json.NewDecoder(r.Body).Decode(&tenant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err := h.s.CreateTenant(&tenant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h TenantHandler) TenantByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant, err := h.s.FindTenantByName(vars["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(tenant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		// TODO: implement
	case http.MethodDelete:
		if err := h.s.DeleteTenant(tenant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
