package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x < y {
		var sum1 int
		for i := x; i <= y; i++ {
			if i%13 != 0 {
				sum1 += i
			}
		}
		fmt.Println(sum1)
	}
	if y < x {
		var sum2 int
		for i := y; i <= x; i++ {
			if i%13 != 0 {
				sum2 += i
			}
		}
		fmt.Println(sum2)
	}
}
