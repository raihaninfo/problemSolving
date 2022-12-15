package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Scan(&x, &y)

		if y == 0 {
			fmt.Println("divisao impossivel")
		} else {
			fmt.Printf("%.1f\n", x/y)

		}

	}
}
