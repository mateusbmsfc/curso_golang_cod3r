package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O maior inteiro é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é", i2)

	// números reais (float32 float64)
	var x float32 = 49.99
	fmt.Println("O valor de x é", reflect.TypeOf(x))
	fmt.Println("O tipo liteial 49.99 é", reflect.TypeOf(49.99)) // float64

	// booleanos
	bo := true
    fmt.Println("O valor de bo é", bo)
    fmt.Println("O valor de bo é com negação", !bo)
    fmt.Println("O tipo literal de bo é", reflect.TypeOf(true))

    // string
    s1 := "Olá"
    fmt.Println("O valor de s1 é", s1)
    fmt.Println("O tipo literal de s1 é", reflect.TypeOf("Olá"))
    fmt.Println("O tipo tamanho da string é", len(s1))

	// string com multriplas linhas
	s2 := `Olá
	meu
	nome
	é
	mateus`
	fmt.Println("O valor de s2 é", s2)
	fmt.Println("O tipo literal de s2 é", reflect.TypeOf(s2))
	fmt.Println("O tipo tamanho da string é", len(s2))

	// char???
	char := 'a'
	fmt.Println("O valor de char é", char)
	fmt.Println("O tipo literal de char é", reflect.TypeOf(char))
}