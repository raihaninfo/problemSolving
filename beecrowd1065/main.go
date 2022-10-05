package main

import "fmt"

func main() {
	var x, y int
	for i := 1; i <= 5; i++ {
		fmt.Scan(&x)
		if x%2 == 0 {
			y++
		}
	}
	fmt.Printf("%v valores pares", y)
}
