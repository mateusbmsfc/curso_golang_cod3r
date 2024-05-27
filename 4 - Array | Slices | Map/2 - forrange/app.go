package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	for i, numero := range numeros {
		fmt.Println(i+1, "-", numero)
	}
}
