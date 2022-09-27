package main

import "fmt"

func main() {
	var n float32
	var s int
	for i := 1; i <= 6; i++ {
		fmt.Scan(&n)
		if n > 0 {
			s = s + 1
		}
	}
	fmt.Printf("%v valores positivos\n", s)
}
