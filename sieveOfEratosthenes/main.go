package main

import "fmt"

func SieveOfEratosthenes(n int) {

	size := n + 1
	prime := make([]bool, size)

	for i := 0; i < size; i++ {
		prime[i] = false
	}

	for i := 2; i*i <= size; i++ {
		if prime[i] == false {
			for j := i * 2; j <= size; j += i {
				prime[j] = true // cross
			}
		}

	}

	// print prime numbers
	fmt.Printf("Primes less than %d : ", n)
	for i := 2; i <= n; i++ {
		if prime[i] == false {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
func main() {
	SieveOfEratosthenes(30)
}
