package solutions

import (
	"encoding/json"
	"fmt"
	"log"
)

func MakeJson() {
	// user input
	var name, address string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your address:")
	fmt.Scanln(&address)

	// Creating a map and adding the name and address to the map using the keys “name” and “address”
	personData := map[string]string{
		"name":    name,
		"address": address,
	}

	//Using Marshal() to create a JSON object from the map
	jsonData, err := json.Marshal(personData)

	//Printing json object
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(jsonData))
	}

}
