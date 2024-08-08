package solutions

import (
	"fmt"
	"math"
	"strconv"
)

// this func checks if given input is valid or not
func isFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// trauncates digits to the right of the decimal place
func TruncateNumber() {
	var input string

	fmt.Print("Enter a floating point number: ")
	fmt.Scanln(&input)

	if isFloat(input) {
		num, _ := strconv.ParseFloat(input, 64)
		truncated := int(math.Trunc(num))
		fmt.Println("Truncated integer:", truncated)
	} else {
		fmt.Println("Invalid input!!!")
	}

}
