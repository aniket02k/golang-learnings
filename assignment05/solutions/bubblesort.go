package solutions

import "fmt"

// Bubble sort algorithm implementation
func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if slice[j] > slice[j+1] {
				swap(&slice, j)
			}
		}
	}
	printSortedSlice(slice)
}

//Function to print slice elements
func printSortedSlice(sortedSlice []int) {
	fmt.Println("Sorted slice elements:")
	for _, val := range sortedSlice {
		fmt.Printf("%d ", val)
	}
}

//swpa function swaps the position of two adjacent elements in the slice
func swap(slice *[]int, i int) {
	temp := (*slice)[i]
	//fmt.Println(temp)
	(*slice)[i] = (*slice)[i+1]
	//fmt.Println((*slice)[i+1])
	(*slice)[i+1] = temp
	//fmt.Println((*slice)[i])

}
