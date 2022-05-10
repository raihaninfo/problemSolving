package main

import (
	"fmt"
	"strings"
)

// Input: address = "1.1.1.1"
// Output: "1[.]1[.]1[.]1"

func main() {
	var address string = "1.1.1.1"
	fmt.Println(address)

	var replaceAddress string = strings.Replace(address, ".", "[.]", -1)
	fmt.Println(replaceAddress)
}

// func defangIPaddr(address string) string {
// 	return strings.Replace(address, ".", "[.]", -1)
// }
