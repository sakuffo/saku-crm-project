package main

import (
	"encoding/json"
	"fmt"
	"io"
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

// SEEDING DATA START

var customer01 Customer = Customer{1, "Steve Kin", "IT Admin", "skin@admin.org", "677-787-1009", true}
var customer02 Customer = Customer{2, "Nomuka Wil", "SRE", "nwil@admin.org", "647-787-1119", false}
var customer03 Customer = Customer{3, "Kacie Oppa", "IT Admin", "kopp@admin.org", "457-755-1002", false}
var customer04 Customer = Customer{4, "Steve Rich Vin Winkle", "Director of Rich", "vanwin@admin.org", "677-787-1009", true}
var customer05 Customer = Customer{5, "Merp P", "Big Bobba Boss", "merp@admin.org", "677-787-2349", true}

var CUSTOMER_DB = map[string]Customer{
	"1": customer01,
	"2": customer02,
	"3": customer03,
	"4": customer04,
	"5": customer05,
}

// SEEDING DATA END

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, customer := range CUSTOMER_DB {
		json.NewEncoder(w).Encode(customer)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(CUSTOMER_DB[id])
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := CUSTOMER_DB[id]; ok {
		delete(CUSTOMER_DB, id)
		getCustomers(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User Not Found")
	}
}

func buildCustomerAddBeta(cust *Customer, json *map[string]string) {

	var test map[string]string = *json
	cust.Id = len(CUSTOMER_DB) + 10
	cust.Name, _ = test["Name"]
	cust.Email, _ = test["Email"]
	cust.Role, _ = test["Role"]
	cust.PhoneNumber, _ = test["PhoneNumber"]
	cust.Contacted, _ = strconv.ParseBool(test["Contacted"])

}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// How to Add a Customer

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var test map[string]string
	json.Unmarshal(reqBody, &test)
	var tempCon Customer
	buildCustomerAddBeta(&tempCon, &test)
	strIndex := strconv.Itoa(len(CUSTOMER_DB) + 10)

	if _, ok := CUSTOMER_DB[strIndex]; ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User Already exists")
	} else {
		CUSTOMER_DB[strIndex] = tempCon
		getCustomers(w, r)
	}
}

func buildCustomerUpdateBeta(cust *Customer, json *map[string]string) {

	var test map[string]string = *json
	cust.Name, _ = test["Name"]
	cust.Email, _ = test["Email"]
	cust.Role, _ = test["Role"]
	cust.PhoneNumber, _ = test["PhoneNumber"]
	cust.Contacted, _ = strconv.ParseBool(test["Contacted"])

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var test map[string]string
	json.Unmarshal(reqBody, &test)
	strIndex := test["Id"]
	tempCon := CUSTOMER_DB[strIndex]
	buildCustomerUpdateBeta(&tempCon, &test)

	if _, ok := CUSTOMER_DB[strIndex]; ok {
		CUSTOMER_DB[strIndex] = tempCon
		getCustomers(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User Not Found")
	}

}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", updateCustomer).Methods("PATCH")

	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server is starting on port 3300...")
	log.Fatal(http.ListenAndServe(":3300", router))

}
