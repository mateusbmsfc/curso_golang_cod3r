package main

import "fmt"

func main() {
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[78945612] = "John"
	aprovados[98765432] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[98765432])
	delete(aprovados, 98765432)
	fmt.Println(aprovados)
}
