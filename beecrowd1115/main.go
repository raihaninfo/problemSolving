package main

import "fmt"

func main() {
	var a, b int

	for {
		fmt.Scan(&a, &b)
		if a == 0 || b == 0 {
			break
		} else if a > 0 && b > 0 {
			fmt.Println("primeiro")
		} else if a > 0 && b < 0 {
			fmt.Println("quarto")
		} else if a < 0 && b < 0 {
			fmt.Println("terceiro")
		} else if a < 0 && b > 0 {
			fmt.Println("segundo")
		}
	}
}
