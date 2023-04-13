package jobs

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
)

type Job struct {
	IDJOB int    `json:"idjob"`
	Job   string `json:"job"`
}

var db *sql.DB
var err error

func GetJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db, err = sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var jobs []Job

	result, err := db.Query("SELECT idjob, job FROM jobs ORDER BY `job`")

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var job Job

		err := result.Scan(&job.IDJOB, &job.Job)

		if err != nil {
			panic(err.Error())
		}

		jobs = append(jobs, job)
	}

	json.NewEncoder(w).Encode(jobs)
}
