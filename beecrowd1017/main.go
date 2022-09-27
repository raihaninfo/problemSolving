package main

import "fmt"

func main() {
	var a, b, res float32
	fmt.Scan(&a, &b)
	res = float32(a * b / 12)
	fmt.Printf("%.3f\n", res)
}
