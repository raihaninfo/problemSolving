package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	x := MairAb(a, b)
	res := MairAb(x, c)
	fmt.Printf("%v eh o maior\n", res)
}

func MairAb(a, b int) int {
	var up int = a + b + abs(a-b)
	var down int = 2
	return up / down
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}
