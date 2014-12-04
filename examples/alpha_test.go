package main

import (
	"fmt"
)

// An example in order to pass all the tests during the build.
func Example() {
	hello()
	// output:
	// Hello World!
}

func hello() {
	fmt.Println("Hello World!")
}
