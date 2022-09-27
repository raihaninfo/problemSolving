package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(time.Month(n))
}
