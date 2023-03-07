package main

import (
    "fmt"
)

func main() {
    var n, w int
    fmt.Scan(&n, &w)

    wei := make([]int, n)
    val := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&wei[i])
    }
    for i := 0; i < n; i++ {
        fmt.Scan(&val[i])
    }

    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, w+1)
    }

    for i := 1; i <= n; i++ {
        for j := 0; j <= w; j++ {
            dp[i][j] = dp[i-1][j]
            if j-wei[i-1] >= 0 {
                dp[i][j] = maximum(dp[i][j], dp[i-1][j-wei[i-1]]+val[i-1])
            }
        }
    }

    fmt.Println(dp[n][w])
}

func maximum(a, b int) int {
    if a > b {
        return a
    }
    return b
}
