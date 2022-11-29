package main

import "fmt"

func main() {
	var a, b int
	for {
		fmt.Scan(&a, &b)
		if a == b {
			break
		} else if a > b {
			fmt.Println("Decrescente")
		} else {
			fmt.Println("Crescente")
		}

	}
}
