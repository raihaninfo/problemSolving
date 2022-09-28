package main

import "fmt"

func main() {
	var pc, q int
	fmt.Scan(&pc, &q)
	var res float32

	if pc == 1 {
		res = 4.00 * float32(q)
	} else if pc == 2 {
		res = 4.50 * float32(q)
	} else if pc == 3 {
		res = 5.00 * float32(q)
	} else if pc == 4 {
		res = 2.00 * float32(q)
	} else if pc == 5 {
		res = 1.50 * float32(q)
	}
	fmt.Printf("Total: R$ %.2f\n", res)

}
