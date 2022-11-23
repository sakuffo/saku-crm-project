package main

import "fmt"

type Customer struct {
	id          int
	name        string
	role        string
	email       string
	phoneNumber string
	contacted   bool
}

func getCustomerInfo(tempQueue map[int]Customer) {
	for _, customer := range tempQueue {
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

func main() {

	customer01 := Customer{1, "Steve Kin", "IT Admin", "skin@admin.org", "677-787-1009", true}
	customer02 := Customer{2, "Nomuka Wil", "SRE", "nwil@admin.org", "647-787-1119", false}
	customer03 := Customer{3, "Kacie Oppa", "IT Admin", "kopp@admin.org", "457-755-1002", false}
	customer04 := Customer{4, "Steve Rich Vin Winkle", "Director of Rich", "vanwin@admin.org", "677-787-1009", true}
	customer05 := Customer{5, "Merp P", "Big Bobba Boss", "merp@admin.org", "677-787-2349", true}

	tempQueue := map[int]Customer{
		1: customer01,
		2: customer02,
		3: customer03,
		4: customer04,
		5: customer05,
	}

	startAppUI()
	getCustomerInfo(tempQueue)

}
