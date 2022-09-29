package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	for i:=x; i<x+12;i++{
		if i%2==1{
			fmt.Println(i)
		}
	}
}