package main

import (
	"fmt"
)

func main() {
	var A, B, C, D, Dif int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)
	Dif = ((A * B) - (C * D))
	fmt.Printf("DIFERENCA = %d\n", Dif)

}
