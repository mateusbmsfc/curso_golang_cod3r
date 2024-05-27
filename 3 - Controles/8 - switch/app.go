package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64, float32:
		return "real"
	case func():
		return "func"
	default:
		return "desconhecido"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("teste"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
