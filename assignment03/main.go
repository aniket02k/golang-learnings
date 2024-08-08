package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	name := getInput("Enter your name:")
	address := getInput("Enter your address:")

	// Create a map to hold the data
	data := make(map[string]string)
	data["name"] = name
	data["address"] = address

	// Marshal the map into a JSON object
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	// Print the JSON object
	fmt.Println("JSON object:", string(jsonData))

}

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
