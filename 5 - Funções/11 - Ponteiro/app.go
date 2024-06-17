package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 5
    inc1(n)
    fmt.Println(n)

	inc2(&n)
	fmt.Println(n)
}