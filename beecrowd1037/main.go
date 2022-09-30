package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	if x >= 0 && x <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if x > 25 && x <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if x > 50 && x <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else if x > 75 && x <= 100 {
		fmt.Println("Intervalo (75,100]")
	} else {
		fmt.Println("Fora de intervalo")
	}
}
