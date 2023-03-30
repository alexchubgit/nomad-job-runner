package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"api/routes/bash"
	"api/routes/companies"
	"api/routes/jobs"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Print("File .env not found")
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/companies", companies.GetCompanies).Methods("GET")
	router.HandleFunc("/jobs", jobs.GetJobs).Methods("GET")
	router.HandleFunc("/bash", bash.GetEcho).Methods("GET")
	router.HandleFunc("/validate", bash.Validate).Methods("GET")

	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)
}
