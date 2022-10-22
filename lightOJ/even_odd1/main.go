package main

import (
	"fmt"
)

func main() {
	var t, v int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&v)
		if v%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
