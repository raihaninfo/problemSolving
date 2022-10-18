package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var j int
	s := strconv.Itoa(n)

	num := []string{}

	for i := len(s) - 1; i >= 0; i-- {
		if j == 3 {
			num = append(num, ",")
			num = append(num, string(s[i]))
			j = 0
		} else {
			num = append(num, string(s[i]))
		}
		j++

	}
	newNum := []string{}
	for i := len(num) - 1; i >= 0; i-- {
		newNum = append(newNum, num[i])
	}
	fmt.Println(strings.Join(newNum, ""))

}
