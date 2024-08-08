/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GrowingSlice() {

	//empty integer slice of size (length) 3
	intSlice := make([]int, 0, 3)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an integer or enter 'X'/'x' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//Exits if user enters 'x' or 'X'
		if strings.ToLower(input) == "x" {
			fmt.Println("Quitting!!")
			break
		}

		//converted string input to the integer value
		num, err := strconv.Atoi(input)
		//validated input
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}

		//appended number entered by user to a slice
		intSlice = append(intSlice, num)
		// sorted slice in increasing order
		sort.Ints(intSlice)
		fmt.Println("Slice elements in sorted order:", intSlice)
	}

}
