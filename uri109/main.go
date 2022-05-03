package main

import "fmt"

func main() {
	var name string
	var salary, sales float64

	fmt.Scan(&name)
	fmt.Scan(&salary)
	fmt.Scan(&sales)

	result := (sales / 100) * 15
	salary = salary + result
	fmt.Printf("TOTAL = R$ %.2f\n", salary)

}
