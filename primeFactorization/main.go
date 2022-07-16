package main

import "fmt"

// Implement prime factorization algorithm
// Input: number
// Output: prime factors
func primeFactorization(number int) []int {
	var primeFactors []int
	// Print the number of 2s that divide number
	for number%2 == 0 {
		primeFactors = append(primeFactors, 2)
		number = number / 2
	}
	for i := 3; i <= number; i = i + 2 {
		// While i divides n, divide n
		for number%i == 0 {
			primeFactors = append(primeFactors, i)
			number = number / i
		}
	}
	if number > 2 {
		primeFactors = append(primeFactors, number)
	}
	return primeFactors
}
func main() {
	fmt.Println(primeFactorization(315)) // [3, 3, 5, 7]
}
