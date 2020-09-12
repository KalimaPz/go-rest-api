package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the My HomePage!")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getPersonInfo", getAllPersonInfo)
	http.ListenAndServe(":8080", nil)
}

func getAllPersonInfo(w http.ResponseWriter, r *http.Request) {
	myPerson := person{
		Firstname: "Donnukrit",
		Lastname:  "Satirakul",
		Code:      1993,
		Phone:     "091-053-3948",
	}
	json.NewEncoder(w).Encode(myPerson) // (1)
}

func main() {
	handleRequest()
}
