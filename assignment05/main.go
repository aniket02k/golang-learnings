package main

import (
	"assignment05/solutions"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter up to 10 integers separated by spaces: ")
	// read user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// split user input on the basis of single whitespace
	numbers := strings.Split(input, " ")
	if len(numbers) > 10 {
		fmt.Println("You've entered more than 10 numbers.Please enter exact 10 numnbers")
		return
	}
	if len(numbers) < 10 {
		fmt.Println("You've entered less than 10 numbers.Please enter exact 10 numnbers")
		return
	}

	//Created slice of integers to store numbers entered by user
	var intNumbers []int
	//Converting each string to integer value and appending to slice
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error occured while converting input to integer value:", err)
			return
		}
		intNumbers = append(intNumbers, num)
	}

	solutions.BubbleSort(intNumbers)
}
