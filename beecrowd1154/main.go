package main

import (
	"fmt"
)

func main() {
	var age, sum, count int
	for {
		fmt.Scan(&age)
		if age < 0 {
			break
		}
		sum += age
		count++

	}

	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Printf("%.2f\n", average)
	}
}
