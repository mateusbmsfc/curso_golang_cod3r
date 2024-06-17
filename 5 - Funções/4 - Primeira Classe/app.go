package main

var soma = func(a, b int) int {
	return a + b
}

func main() {
    resultado := soma(1, 2)
    println(resultado)

	sub := func(a, b int) int {
		return a - b
	}

	resultado = sub(1, 2)
	println(resultado)
}