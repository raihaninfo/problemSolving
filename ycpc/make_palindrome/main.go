package main

import "fmt"

func main() {

	var S string
	fmt.Scan(&S)

	var i, j int
	var count int
	for i, j = 0, len(S)-1; i < j; i, j = i+1, j-1 {
		if S[i] != S[j] {
			count++
		}
	}

	fmt.Println(count)

}
