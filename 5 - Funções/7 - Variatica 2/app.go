package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, nome := range aprovados {
		fmt.Println(i+1, nome)
    }
}

func main() {
    imprimirAprovados("Ana", "Bia", "Carlos", "Daniel")
}