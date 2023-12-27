package main

import (
	"fmt"
	"log"
	"net/http"

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

	r := mux.NewRouter()
	apartmentHandler := handlers.NewApartmentHandler(db)

	r.HandleFunc("/apartments/", apartmentHandler.ApartmentList)
	r.HandleFunc("/apartments/{number}", apartmentHandler.ApartmentByNumber)

	log.Fatal(http.ListenAndServe(":8080", r))
}
