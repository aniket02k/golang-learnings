package main

import (
	"assignment06/solution"
	"fmt"
)

func main() {
	var a, vo, so, t float32

	// User input for accelaration, veloccity and displacement
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vo)

	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&so)

	// Create the displacement function with the given parameters
	fn := solution.GenDisplaceFn(a, vo, so)

	// Prompt user for time
	fmt.Print("Enter time (t): ")
	fmt.Scan(&t)

	// Calculate and print the displacement after the given time
	displacement := fn(t)
	fmt.Printf("The displacement after %.2f seconds is: %.2f\n", t, displacement)

}
