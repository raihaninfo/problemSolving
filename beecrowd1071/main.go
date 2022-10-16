package main

import "fmt"

func main(){
	var x, y, sum int
	fmt.Scan(&x, &y)
	var min, max int 

	if x<y{
		min = x
		max= y
	}else{
		min = y
		max= x
	}

	for i:=min+1; i<max; i++{
		if i%2!=0{
			sum+=i
		}
	}
	fmt.Println(sum)
}