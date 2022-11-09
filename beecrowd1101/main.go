// https://www.beecrowd.com.br/judge/en/problems/view/1101
package main

import (
	"fmt"
)

func main() {
	var m, n int
	for {
		fmt.Scan(&m, &n)
		if m <= 0 || n <= 0 {
			break
		}
		if m > n {
			m, n = n, m
		}
		sum := 0
		for i := m; i <= n; i++ {
			fmt.Printf("%d ", i)
			sum += i
		}
		fmt.Printf("Sum=%d", sum)
		fmt.Println()

	}

}
