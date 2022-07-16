package main

import (
	"fmt"
	"github.com/Property-Finder-Patika/week-3-homework-1-erdemcemal/inanc-gumus/fix-the-memory-leak/api"
	"io/ioutil"
)

func main() {
	// reports the initial memory usage

	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := api.Read()

	// ------------------------------------------------------

	last10 := make([]int, 10)
	copy(last10, millions[len(millions)-10:])

	// Make the millions slice lose reference to its backing array
	// so that its backing array can be cleaned up from memory.
	millions = last10

	fmt.Printf("\nLast 10 elements: %d\n\n", last10)

	// ------------------------------------------------------

	api.Report()

	// don't worry about this code yet.
	fmt.Fprintln(ioutil.Discard, millions[0])
}
