package main

import (
	"assignment07/solution"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var animals = [...]solution.Animal{
		solution.Animal{"cow", "grass", "walk", "moo"},
		solution.Animal{"bird", "worms", "fly", "peep"},
		solution.Animal{"snake", "mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please provide an animal and a request separated by a space.")
			continue
		}
		animalName := parts[0]
		request := parts[1]
		for _, a := range animals {
			if a.Name == animalName {
				switch request {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			}

		}

	}

}
