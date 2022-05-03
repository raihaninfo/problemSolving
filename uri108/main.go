package main

import "fmt"

func main() {
	var FuncNum, Horas int
	var valor, salario float64

	fmt.Scan(&FuncNum)
	fmt.Scan(&Horas)
	fmt.Scan(&valor)

	salario = float64(Horas) * valor

	fmt.Printf("NUMBER = %d\n", FuncNum)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
