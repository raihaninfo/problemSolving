package main

import "fmt"

func main() {
	var x, n, p, o, e int

	for i := 1; i <= 5; i++ {
		// fmt.Println("lsjdf", i)
		fmt.Scan(&x)
		if x%2 == 0 {
			e++
		} else {
			o++
		}
		if x > 0 {
			p++
		}
		if x < 0 {
			n++
		}
	}
	fmt.Printf("%v valor(es) par(es)\n%v valor(es) impar(es)\n%v valor(es) positivo(s)\n%v valor(es) negativo(s)\n", e, o, p, n)
}
