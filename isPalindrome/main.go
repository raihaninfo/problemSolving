package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	arr := []int{}

	for i := 0; i < len(s); i++ {
		arr = append(arr, int(s[i]))
		// arr[i] = int(s[i])
	}

	arr2 := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		arr2 = append(arr2, int(arr[i]))
		// arr2[i] = int(arr[i])
	}
	var result bool
	// fmt.Println(arr==arr2)
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr2[i] {
			result = true
			continue
		} else {
			result = false
			break
		}
	}
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
