package main

import "fmt"

func main() {
	funcs := map[string]float64{
		"Maria": 15000.00,
		"John":  10000.00,
		"Ana":   90000.50,
	}
	fmt.Println(funcs)

	funcs["Rafael"] = 1350.00
	fmt.Println(funcs)

	delete(funcs, "Inexistente")

	for nome, salario := range funcs {
		fmt.Printf("%s (R$ %.2f)\n", nome, salario)
	}
}
