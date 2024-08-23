package solution

import "fmt"

// Animal struct to hold information about each animal
type Animal struct {
	Name       string
	Food       string
	Locomotion string
	Noise      string
}

// prints the food that the animal eats

func (a Animal) Eat() {
	fmt.Println(a.Name, "eats:", a.Food)
}

// prints the locomotion method of the animal
func (a Animal) Move() {
	fmt.Println(a.Name, "moves by:", a.Locomotion)
}

// prints the sound the animal makes
func (a Animal) Speak() {
	fmt.Println(a.Name, "speaks with:", a.Noise)
}
