package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxSum(n, k int, arr []int) int {
	var maxSum, currentSum int

	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}

	maxSum = currentSum

	for i := k; i < n; i++ {
		currentSum += arr[i] - arr[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	var t, n, k int
	fmt.Scan(&t)

	reader := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n, &k)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &arr[j])
		}
		fmt.Fprintln(write, maxSum(n, k, arr))
	}

	write.Flush()
}
