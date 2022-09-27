package main

import "fmt"

func main() {
	var dist, fuel, avg float64
	fmt.Scan(&dist, &fuel)
	avg = dist / fuel
	fmt.Printf("%.3f km/l\n", avg)
}
