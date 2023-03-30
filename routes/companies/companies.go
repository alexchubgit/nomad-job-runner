package companies

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Company struct {
	IDCOMPANY int    `json:"idcompany"`
	Company   string `json:"company"`
}

var db *sql.DB
var err error

func GetCompanies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var companies []Company

	result, err := db.Query("SELECT idcompany, company FROM companies ORDER BY `company`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var company Company

		err := result.Scan(&company.IDCOMPANY, &company.Company)

		if err != nil {
			panic(err.Error())
		}

		companies = append(companies, company)
	}

	json.NewEncoder(w).Encode(companies)
}
