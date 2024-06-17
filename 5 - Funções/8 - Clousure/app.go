package main

import "fmt"

func closure() func() {
	x := 0
	var funcao = func() {
		fmt.Println(x)
    }
    return funcao
}

func main() {
    x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}