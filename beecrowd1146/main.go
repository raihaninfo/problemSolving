package main

import "fmt"

func main() {
	var x int

	for {
		fmt.Scan(&x)

		if x > 0 {
			for i := 1; i < x; i++ {
				fmt.Printf("%d ", i)
			}
		}
		if x != 0 {
			fmt.Printf("%d\n", x)
		}
		if x == 0 {
			break
		}
	}
}
