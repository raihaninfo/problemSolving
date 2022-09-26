package main

import "fmt"

func main() {
	var a, b int
	var c, res float32
	fmt.Scan(&a, &b, &c)
	res = float32(b) * c
	fmt.Scan(&a, &b, &c)
	res += float32(b) * c

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", res)
}
