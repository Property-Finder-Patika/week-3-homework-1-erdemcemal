package main

import "fmt"

// Write a function that rotates given slice by given number of steps to the right or to the left.
// Input: slice, number of steps
// Output: rotated slice
func rotate(slice []int, steps int, isLeft bool) []int {
	if steps == 0 || len(slice) == 0 {
		return slice
	}
	if isLeft {
		return leftRotate(slice, steps)
	}
	return rightRotate(slice, steps)
}

// leftRotate is a helper function for rotate a slices on left direction
func leftRotate(slice []int, steps int) []int {
	return append(slice[steps:], slice[:steps]...)
}

// rightRotate is a helper function for rotate a slices on right direction
func rightRotate(slice []int, steps int) []int {
	return append(slice[len(slice)-steps:], slice[:len(slice)-steps]...)
}
func main() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5}, 2, true))  // [3, 4, 5, 1, 2]
	fmt.Println(rotate([]int{1, 2, 3, 4, 5}, 2, false)) // [4, 5, 1, 2, 3]
}
