package main

import "fmt"

func main(){
	var t int
	var s, c string

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&s, &c)
		for j := 0; j < len(s); j++ {
			if s[j] != c[0] {
				fmt.Printf("%v", string(s[j]))
			}
		}
		fmt.Println()
	}

}