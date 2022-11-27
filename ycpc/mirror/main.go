package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i == 1 || j == 1 {
				fmt.Print(j + i - 1)
			} else if j == N {
				fmt.Print(N - i + 1)
			} else if i == N {
				fmt.Print(N - j + 1)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
