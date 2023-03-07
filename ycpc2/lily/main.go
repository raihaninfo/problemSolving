package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    t, _ := strconv.Atoi(scanner.Text())

    for i := 0; i < t; i++ {
        scanner.Scan()
        n, _ := strconv.Atoi(scanner.Text())

        scanner.Scan()
        arr := make([]int, n)
        nums := strings.Split(scanner.Text(), " ")
        for j := 0; j < n; j++ {
            arr[j], _ = strconv.Atoi(nums[j])
        }

        for j := 0; j < n-1; j++ {
            for k := j + 1; k < n; k++ {
                if arr[j] < arr[k] {
                    arr[j], arr[k] = arr[k], arr[j]
                }
            }
        }

        even, odd := -1, -1
        for j := 0; j < n; j++ {
            if arr[j]%2 == 0 {
                even = arr[j]
            } else {
                odd = arr[j]
            }
            if even != -1 && odd != -1 {
                fmt.Println("YES")
                break
            }
        }
        if even == -1 || odd == -1 {
            fmt.Println("NO")
        }
    }
}
