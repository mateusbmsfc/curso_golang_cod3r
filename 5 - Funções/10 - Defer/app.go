package main

import "fmt"

func obterValorAprovado(num int) int {
	defer fmt.Println("FIM!")
	if num >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} else {
		fmt.Println("Valor baixo...")
        return num
	}
}

func main() {
    valor := obterValorAprovado(10000)
    fmt.Println(valor)
}