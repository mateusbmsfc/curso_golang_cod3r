package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)

	// apenas postfix
	x++
	y--
	fmt.Println(x, y)

	// Vai dar erro abaixo
	// fmt.Println( x == y--)
}