package main

import "fmt"

func main() {
	var A, B, min, max int
	fmt.Scan(&A, &B)

	// min, max
	if A > B {
		min = B
		max = A
	} else {
		min = A
		max = B
	}

	// difference
	fmt.Println(max - min)

}
