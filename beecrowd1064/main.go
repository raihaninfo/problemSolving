package main

import "fmt"

func main(){
	var x, y,avg float32
	var a int

	for i := 0; i < 6; i++ {
		fmt.Scan(&x)
		if x>0{
			y+=x
			a++
		}
	}
	avg= y/float32(a)
	fmt.Printf("%v valores positivos\n%.1f\n", a, avg)
}