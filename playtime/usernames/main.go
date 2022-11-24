package main

import (
	"fmt"
	"strings"
)

func main() {
	usernames := map[string]string{
		"Sammy": "sammy-shark",
		"Jamie": "mantisshrimp54",
	}

	for {
		// Ask the user for input, a name
		fmt.Println("Enter a Name:")
		var name string
		_, err := fmt.Scanln(&name)
		if err != nil {
			panic(err)
		}

		// Trim the white
		name = strings.TrimSpace(name)

		if u, ok := usernames[name]; ok {
			fmt.Printf("%q is the username of %q\n", u, name)
			continue
		}

		fmt.Printf("I don't have %v's username, what is it?\n", name)

		var username string
		_, err = fmt.Scanln(&username)
		if err != nil {
			panic(err)
		}

		username = strings.TrimSpace(username)
		usernames[name] = username
		fmt.Println("Data updated.")
		fmt.Println(usernames)
	}
}
