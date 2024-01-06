package handlers

import (
	"benttreeGo/pkg/models"
	"benttreeGo/pkg/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TenantHandler struct {
	s services.DatabaseService
}

func NewTenantHandler(s services.DatabaseService) *TenantHandler {
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
			return
		}
		if err = json.NewEncoder(w).Encode(tenants); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
		if err := h.s.PutTenant(*tenant); err != nil {
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
			if err := h.s.PatchTenant(vars["name"], k, v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	case http.MethodDelete:
		if err := h.s.DeleteTenant(tenant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// eventually maybe add PUT and PATCH methods
func (h TenantHandler) TenantsByApartmentNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenants, err := h.s.FindTenantsByApartmentNumber(vars["apartment_number"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(tenants); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
