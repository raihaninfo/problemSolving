package main

import (
    "fmt"
    "strconv"
)

func applyString(inputString string) string {
    result := ""
    i := 0
    for i < len(inputString) {
        statement := ""
        numberString := ""
        
        for i < len(inputString) && inputString[i] >= '0' && inputString[i] <= '9' {
            numberString += string(inputString[i])
            i++
        }
        
        for i < len(inputString) && inputString[i] >= 'a' && inputString[i] <= 'z' {
            statement += string(inputString[i])
            i++
        }
        
        N, _ := strconv.Atoi(numberString)
        for j := 0; j < N; j++ {
            result += statement
        }
        
        if i < len(inputString) && (inputString[i] == '+' || inputString[i] == '-' || inputString[i] == '*' || inputString[i] == '/') {
            result += string(inputString[i])
            i++
        }
    }
    return result
}

func main() {
    var inputString string
    fmt.Scan(&inputString)
    fmt.Println(applyString(inputString))
}
