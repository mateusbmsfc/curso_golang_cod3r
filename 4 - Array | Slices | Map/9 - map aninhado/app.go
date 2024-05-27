package main

import "fmt"

func main() {
	funcsPorLetrs := map[string]map[string]float32{
		"G": {
			"Gabriela": 15000.00,
			"Gabriel":  10000.00,
			"Gaby":     90000.50,
		},
		"J": {
			"José": 1350.00,
		},
		"M": {
			"Mateus": 25000.00,
		},
	}
	fmt.Println(funcsPorLetrs)

	delete(funcsPorLetrs, "J")

	for letra, funcs := range funcsPorLetrs {
		fmt.Println(letra, funcs)
	}
}
