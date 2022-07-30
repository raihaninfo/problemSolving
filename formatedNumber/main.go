package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := strconv.Itoa(n)

	num := []int{}

	for i := len(s) - 1; i >= 0; i-- {
		num = append(num, int(s[i]-'0'))
	}
	fmt.Println(num)

}
