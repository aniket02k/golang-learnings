package solutions

import (
	"fmt"
	"strings"
)

func FindIAN() {
	var inputStr string
	fmt.Println("Enter a string:")
	fmt.Scanln(&inputStr)

	//convert string to lower case to avoid case-sensitivity.
	inputStr = strings.ToLower(inputStr)
	//logic to check if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’
	if strings.HasPrefix(inputStr, "i") && strings.HasSuffix(inputStr, "n") && strings.Contains(inputStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	//fmt.Println(inputStr)

}
