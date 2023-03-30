package bash

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type Company struct {
	IDCOMPANY int    `json:"idcompany"`
	Company   string `json:"company"`
}

func GetEcho(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	app := "echo"

	arg0 := "-e"
	arg1 := "Hello world"
	arg2 := "\n\tfrom"
	arg3 := "golang"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

	json.NewEncoder(w).Encode(string(stdout))
}

func Validate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	app := "nomad"

	arg0 := "job"
	arg1 := "validate"
	arg2 := "./files/docs.job"

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

	json.NewEncoder(w).Encode(string(stdout))
}
