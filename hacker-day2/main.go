package main

import "fmt"

func main() {
	//main code
	var n int
	fmt.Scan(&n)

	for i := 1; i <=10; i++ {
		// result:= 
		fmt.Println(n, "x", i, "=", n*i)
	}
}
