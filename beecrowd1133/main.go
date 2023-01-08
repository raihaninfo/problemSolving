package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > y {
		for i := y+1; i < x; i++ {
			if i%5 == 2 || i%5 == 3 {
				fmt.Println(i)
			}
		}
	}

	if x < y {
		for i := x+1; i < y; i++ {
			if i%5 == 2 || i%5 == 3 {
				fmt.Println(i)
			}
		}
	}
}
