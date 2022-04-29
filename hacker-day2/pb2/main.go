package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2==1 || n>5&&n<21{
		fmt.Println("Weird")
	}else{
		fmt.Println("Not Weird")
	}
}
