package main

import "fmt"

func main() {

// 234234
	var a int
	fmt.Scan(&a)

	// hour
	var b int = a / 3600

	a = a - b*3600

	// minute
	var c int = a / 60
	a = a - c*60

	// secund
	var d int = a

	fmt.Printf("%v:%v:%v\n", b, c, d)
}
