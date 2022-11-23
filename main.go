package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	id          int
	name        string
	role        string
	email       string
	phoneNumber string
	contacted   bool
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

// UI Functions

func printCustomerInfo() {
	for _, customer := range CUSTOMER_DB {
		fmt.Println("----------------------------")
		fmt.Printf("Customer id: %d\n", customer.id)
		fmt.Printf("Customer name: %s\n", customer.name)
		fmt.Printf("Customer role: %s\n", customer.role)
		fmt.Printf("Customer email: %s\n", customer.email)
		fmt.Printf("Customer phoneNumber: %s\n", customer.phoneNumber)
		fmt.Printf("Customer contacted: %t\n", customer.contacted)
		fmt.Println("----------------------------")
	}
}

func startAppUI() {
	fmt.Printf("----------------------------\n")
	fmt.Println("CRM Backend Start")
	fmt.Printf("----------------------------\n\n\n")
}

// CRUD Functions

// func getAllCustomers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	customerDB_json, err := json.Marshal(CUSTOMER_DB)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(CUSTOMER_DB)

// }

func seedCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	for _, customer := range CUSTOMER_DB {
		fmt.Fprintf(
			w,
			"<p>ID: %d\t| NAME: %s\t\t\t| ROLE: %s\t\t\t| EMAIL: %s\t\t\t| PNUMBER: %s\t\t\t| CONTACTED: %t</p>",
			customer.id, customer.name, customer.role, customer.email, customer.phoneNumber, customer.contacted)
	}

	w.WriteHeader(http.StatusOK)
	// iterate over cityPoulations and print them within an h2
}

func main() {

	// start application banner
	startAppUI()

	//  helper function to enumerate over mock initial
	// customerDB
	printCustomerInfo()

	// 21:31 nov 22, starting to set up routes for API
	router := mux.NewRouter().StrictSlash(true)

	// Handlers
	router.HandleFunc("/", seedCheck).Methods("GET")
	// router.HandleFunc("/customers", getAllCustomers).Methods("GET") I have an issue encoding the map as json right now.

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3300", router))

}
