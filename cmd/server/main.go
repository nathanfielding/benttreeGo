package main

import (
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

	file, err := os.Open("schema.sql")
	if err != nil {
		panic(err)
	}
	schema := make([]byte, 1024)
	file.Read(schema)

	_, err = db.Exec(string(schema))
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	apartmentHandler := handlers.NewApartmentHandler(db)

	r.HandleFunc("/apartments/", apartmentHandler.ApartmentList)
	r.HandleFunc("/apartments/{number}", apartmentHandler.ApartmentByNumber)

	log.Fatal(http.ListenAndServe(":8080", r))
}
