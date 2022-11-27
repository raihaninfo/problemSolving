package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	// var S1, S2 string

	if A > B {
		fmt.Println("Argentina")
	} else if A == B {
		

	} else {
		fmt.Println("Brazil")
	}

}

// create a function compare that two numbers in string format
// func compare(a, b string) (int, int) {
// 	var aa, bb int
// 	for i := 0; i < len(a); i++ {
// 		if a[i] == '1' {
// 			aa++
// 		}
// 		if b[i] == '1' {
// 			bb++
// 		}

// 	}

// 	return aa, bb

// }
