package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)

	for i:=1;i<=10000; i++{
		if i%x==2{
			fmt.Println(i)
		}
	}
}