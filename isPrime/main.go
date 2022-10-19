package main

import "fmt"

func main(){
	var n int
	var count int =0
	fmt.Scan(&n)

	for i:= 2; i<n; i++{
		if n%i==0{
			count ++
		}
	}
	if count==0{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}