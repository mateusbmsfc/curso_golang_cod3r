package main

import "fmt"

func main() {
	i := 1
	
	// GO NÃO TEM ARITMETICA DE PONTEIRO
	var p *int = nil
	p = &i
	*p++
	fmt.Println(i)
	i++
	fmt.Println(i)
	fmt.Println(p)
}