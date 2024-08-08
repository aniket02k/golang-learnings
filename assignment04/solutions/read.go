package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

var people []Person

func ReadData() {
	var filename string
	fmt.Println("Enter a file name from which you want to read a person's data")
	fmt.Scanln(&filename)
	data, err := os.Open(filename)
	if err != nil {
		fmt.Println("File not Found. Please enter correct file name!")
		return
	}
	defer data.Close()

	//fmt.Printf("%T", data)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		lineInFile := scanner.Text()
		personInfo := strings.Split(lineInFile, " ")

		person := Person{
			fname: personInfo[0],
			lname: personInfo[1],
		}
		people = append(people, person)
	}
	printAllRecords(people)
}

func printAllRecords(peopleList []Person) {
	for _, person := range peopleList {
		fmt.Printf("First Name: %s\t Last Name: %s\n", person.fname, person.lname)
	}
}
