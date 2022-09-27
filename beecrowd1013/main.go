package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var res int
	if a > b && a > c {
		res = a
	} else if b > a && b > c {
		res = b
	} else {
		res = c
	}
	fmt.Printf("%v eh o maior", res)

}
