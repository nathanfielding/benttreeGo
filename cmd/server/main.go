package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"benttreeGo/pkg/handlers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting server...")
	postresqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "postgres", 5432, "benttree_user", "benttree_password", "benttree_db")
	db, err := sqlx.Open("postgres", postresqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	schemaPath := flag.String("path", "/benttreeGo.schema.sql", "Path to schema file")
	flag.Parse()

	schema, err := os.ReadFile(*schemaPath)
	if err != nil {
		panic(err)
	}

	db.MustExec(string(schema))

	r := mux.NewRouter()
	apartmentHandler := handlers.NewApartmentHandler(db)

	r.HandleFunc("/apartments/", apartmentHandler.ApartmentList)
	r.HandleFunc("/apartments/{number}", apartmentHandler.ApartmentByNumber)

	log.Fatal(http.ListenAndServe(":8080", r))
}
