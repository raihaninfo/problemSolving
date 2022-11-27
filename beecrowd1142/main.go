package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d %d %d PUM", i, i+1, i+2)
			i += 4
			fmt.Println()
		}
	}

}
