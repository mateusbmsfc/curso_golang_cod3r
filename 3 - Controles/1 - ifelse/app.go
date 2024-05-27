package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

func main() {
	imprimirResultado(6.9)
	imprimirResultado(2.9)
	imprimirResultado(10.0)
}
