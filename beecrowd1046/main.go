package main

import "fmt"

func main() {
	var st, end, res int
	fmt.Scan(&st, &end)

	if st < end {
		res = end - st
	} else {
		res = 24 - st + end
	}
	fmt.Printf("O JOGO DUROU %v HORA(S)\n", res)
}
