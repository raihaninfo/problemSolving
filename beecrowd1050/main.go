package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	// declare and initialize map
	m := map[int]string{
		61: "Brasilia",
		71: "Salvador",
		11: "Sao Paulo",
		21: "Rio de Janeiro",
		32: "Juiz de Fora",
		19: "Campinas",
		27: "Vitoria",
		31: "Belo Horizonte",
	}

	if len(m[a]) > 0 {
		fmt.Println(m[a])
	} else {
		fmt.Println("DDD nao cadastrado")
	}
}
