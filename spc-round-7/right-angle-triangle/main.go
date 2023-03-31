package main

import "fmt"

func main(){
	var a int
	var x, y, z float64
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&x, &y, &z)
		if x*x + y*y == z*z || y*y + z*z == x*x || z*z + x*x == y*y {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
		
}