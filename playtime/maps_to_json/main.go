package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Customer struct {
	id          int
	Name        string
	Role        string
	email       string
	phoneNumber string
	contacted   bool
}

func main() {
	// R&D on how to pass a map[int]string as json

	// attempting to pass a map with int keys is failing in crm backend
	// trying to see if there is a way to get around this

	// map1 := map[int]string{
	// 	1: "apple",
	// 	2: "banana",
	// }

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

	// CUST_DB2 := map[string]Customer{
	// 	"1": {"Name": "Steve Kin", "Role": "IT Admin"},
	// 	"2": {"Name": "Nomuka Wil", "Role": "SRE"},
	// 	"3": {"Name": "Kacie Oppa", "Role": "IT Admin"},
	// 	"4": {"Name": "Steve Rich Vin Winkle", "Role": "Director of Rich"},
	// 	"5": {"Name": "Merp P", "Role": "Big Bobba Boss"},
	// }

	// jsonStr, err := json.Marshal(map1)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// } else {
	// 	fmt.Println(string(jsonStr))
	// }

	// jsonStr2, err := json.Marshal(CUSTOMER_DB)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// } else {
	// 	fmt.Println(string(jsonStr2))
	// }

	fmt.Println(CUSTOMER_DB)
	testing := f"testomg"
	fmt.Println(strconv.Quote(CUSTOMER_DB))
}
