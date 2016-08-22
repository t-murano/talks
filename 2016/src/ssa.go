package main

import "fmt"

func main() {
	a := 1         // Define: a1 = 1
	b := a * 2     // Define: b1 = a1 * 2   Use: a1
	a = 10         // Define: a2 = 10
	c := a + b     // Define: c1 = a2 + b1  Use: a2, b1
	fmt.Println(c) // Use: c1
}
