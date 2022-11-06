// https://www.beecrowd.com.br/judge/en/problems/view/1074
package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		if x == 0 {
			fmt.Println("NULL")
		} else if x < 0 && x%2 == 0 {
			fmt.Println("EVEN NEGATIVE")
		} else if x > 0 && x%2 == 0 {
			fmt.Println("EVEN POSITIVE")
		} else if x < 0 && x%2 !=0 {
			fmt.Println("ODD NEGATIVE")
		} else if x > 0 && x%2 !=0 {
			fmt.Println("ODD POSITIVE")
		}
	}
}
