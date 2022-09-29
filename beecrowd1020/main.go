package main

import "fmt"

func main() {
	var day int
	fmt.Scan(&day)

	year := day / 365
	month := (day-year*365)/30
	day = day-year*365 - month *30
	fmt.Printf("%v ano(s)\n%v mes(es)\n%v dia(s)\n", year, month, day)

}
