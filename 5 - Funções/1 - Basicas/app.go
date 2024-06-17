package main

import "fmt"

// Função sem parametro e sem retorno
func f1() {
	fmt.Println("F1")
}

// Tem parametro e sem retorno
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// Não tem paramentro e retorna string
func f3() string {
	return "F3"
}

// Tem paramentro e retorna string
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s\n", p1, p2)
}

func f5() (string, string) {
	return "a", "b"
}

func main() {
	f1()
    f2("a", "b")
    
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3, r4)
	fmt.Println(f5())
}