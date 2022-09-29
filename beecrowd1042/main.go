package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	n := []int{a, b, c}

	sort.Ints(n)
	aa := n[0]
	bb := n[1]
	cc := n[2]
	fmt.Printf("%v\n%v\n%v\n\n%v\n%v\n%v\n", aa, bb, cc, a, b, c)
}
