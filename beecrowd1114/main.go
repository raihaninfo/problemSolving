package main

import "fmt"

func main() {
	var pass int
	for {
		fmt.Scan(&pass)

		if pass == 2002 {
			fmt.Println("Acesso Permitido")
			break
		} else {
			fmt.Println("Senha Invalida")
		}
	}

}
