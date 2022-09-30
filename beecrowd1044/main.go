package main

import "fmt"

func main(){
	var a,b int
	fmt.Scan(&a, &b)

	if a%b==0 || b%a==0{
		fmt.Println("Sao Multiplos")
	}else{
		fmt.Println("Nao sao Multiplos")
	}
}