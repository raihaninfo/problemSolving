package main

import "fmt"

func main(){
	var n, highest, pos int
	for i := 1; i <= 100; i++ {
		fmt.Scan(&n)
		if n>highest{
			highest = n
			pos = i
		}
	}
	fmt.Println(highest)
	fmt.Println(pos)
	
}