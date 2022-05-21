package main

import "fmt"

func main() {
	num := []int{2, 4, 5, 6, 7}
	fmt.Println(runningSum(num))
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// input [2,4,5,6,7]
// output [2,6,11,17,24]
