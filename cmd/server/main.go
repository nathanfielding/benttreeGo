package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"benttreeGo/pkg/handlers"
	"benttreeGo/pkg/services"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	schemaPath := flag.String("path", "/benttreeGo.schema.sql", "Path to schema file")
	dbHost := flag.String("host", "localhost", "System to run database on")
	dbPort := flag.Int("port", 5432, "Port to open database connection on")
	flag.Parse()

	postresqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *dbHost, *dbPort, "benttree_user", "benttree_password", "benttree_db")

	db, err := sqlx.Open("postgres", postresqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	schema, err := os.ReadFile(*schemaPath)
	if err != nil {
		panic(err)
	}

	db.MustExec(string(schema))

	r := mux.NewRouter()
	dbService := services.NewDatabaseService(db)

	apartmentHandler := handlers.NewApartmentHandler(dbService)
	tenantHandler := handlers.NewTenantHandler(dbService)
	leaseHandler := handlers.NewLeaseHandler(dbService)

	r.HandleFunc("/apartments/", apartmentHandler.ApartmentList)
	r.HandleFunc("/apartments/number/{number}", apartmentHandler.ApartmentByNumber)

	r.HandleFunc("/tenants/", tenantHandler.TenantList)
	r.HandleFunc("/tenants/name/{name}", tenantHandler.TenantByName)
	r.HandleFunc("/tenants/number/{number}", tenantHandler.TenantsByApartmentNumber)

	r.HandleFunc("/leases/", leaseHandler.LeaseList)
	r.HandleFunc("/leases/name/{name}", leaseHandler.LeaseByName)

	log.Fatal(http.ListenAndServe(":8080", r))
}
