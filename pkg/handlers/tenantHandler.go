package handlers

import (
	"benttreeGo/pkg/models"
	"benttreeGo/pkg/services"
	"encoding/json"
	"net/http"
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
