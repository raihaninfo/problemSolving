package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A > B {
		fmt.Println("Argentina")
	} else if A == B {
		fmt.Println("Draw")
	} else {
		fmt.Println("Brazil")
	}
}
