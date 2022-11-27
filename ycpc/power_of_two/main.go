package main

import "fmt"

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Println(powerOfTwo(N))
	}

}

func powerOfTwo(N int) int {
	if N == 0 {
		return 1
	}

	var i int
	for i = 1; i < N; i *= 2 {
	}

	return i
}
