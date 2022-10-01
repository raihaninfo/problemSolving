package main

import "fmt"

func main() {
	var n, pa, new, p float32
	fmt.Scan(&n)

	if n >= 0 && n <= 400.00 {
		p = 15
		pa = n / 100 * p
		new = n + pa
	} else if n >= 400.01 && n <= 800.00 {
		p = 12
		pa = n / 100 * p
		new = n + pa
	} else if n >= 800.01 && n <= 1200.00 {
		p = 10
		pa = n / 100 * p
		new = n + pa
	} else if n >= 1200.01 && n <= 2000.00 {
		p = 7
		pa = n / 100 * p
		new = n + pa
	} else if n > 2000.00 {
		p = 4
		pa = n / 100 * p
		new = n + pa
	}

	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %v %%\n", new, pa, p)
}
