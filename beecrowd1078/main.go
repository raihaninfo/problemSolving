package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		j := i * n
		fmt.Printf("%v x %v = %v\n", i, n, j)
	
	}
}
