package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id          int
	Name        string
	Role        string
	Email       string
	PhoneNumber string
	Contacted   bool
}

var customer01 Customer = Customer{1, "Steve Kin", "IT Admin", "skin@admin.org", "677-787-1009", true}
var customer02 Customer = Customer{2, "Nomuka Wil", "SRE", "nwil@admin.org", "647-787-1119", false}
var customer03 Customer = Customer{3, "Kacie Oppa", "IT Admin", "kopp@admin.org", "457-755-1002", false}
var customer04 Customer = Customer{4, "Steve Rich Vin Winkle", "Director of Rich", "vanwin@admin.org", "677-787-1009", true}
var customer05 Customer = Customer{5, "Merp P", "Big Bobba Boss", "merp@admin.org", "677-787-2349", true}

var CUSTOMER_DB = map[int]Customer{
	1: customer01,
	2: customer02,
	3: customer03,
	4: customer04,
	5: customer05,
}

// CRUD Functions

func getAllCustJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, customer := range CUSTOMER_DB {
		json.NewEncoder(w).Encode(customer)
	}
}

func getCustJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		fmt.Print("Error")
	}
	json.NewEncoder(w).Encode(CUSTOMER_DB[idNumber])
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	if _, ok := CUSTOMER_DB[id]; ok {
		delete(CUSTOMER_DB, id)
		getAllCustJSON(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User Not Found")
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// How to Add a Customer
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// How to Update a Customer
}

func index(w http.ResponseWriter, r *http.Request) {
	// Overview and Instructions on how to use the API
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>CRM BACKEND API | WORK IN PROGRESS</h1>")
	w.WriteHeader(http.StatusOK)
}

func main() {

	// 21:31 nov 22, starting to set up routes for API
	router := mux.NewRouter().StrictSlash(true)

	// Handlers || CRUD
	router.HandleFunc("/", index).Methods("GET") // Basic Implemented. Copy for instructions left []

	router.HandleFunc("/customers", addCustomer).Methods("POST")     // Not implemented // CREATE
	router.HandleFunc("/customers", getAllCustJSON).Methods("GET")   // READ Implemented. Uses a for loop to render the json
	router.HandleFunc("/customers", updateCustomer).Methods("PATCH") // Not Implemented // UPDATE

	router.HandleFunc("/customers/{id}", getCustJSON).Methods("GET")          // READ 1 Implemented
	router.HandleFunc("/deleteCustomer/{id}", deleteCustomer).Methods("POST") // DELETE Implemented

	fmt.Println("Server is starting on port 3300...")
	log.Fatal(http.ListenAndServe(":3300", router))

}
