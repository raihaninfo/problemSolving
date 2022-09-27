package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2, res float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	res = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("%.4f\n", res)
}
