package main

// NÃO TEM OPERADOR TERNARIO

import "fmt"

func obterResult(nota float64) string {
	if nota >= 7 {
        return "Aprovado"
    } else {
        return "Reprovado"
    }
}

func main() {
    fmt.Println(obterResult(6.9))
    fmt.Println(obterResult(2.9))
}