package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMistrurando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par... ")
		} else {
			fmt.Printf("Impar... ")
		}
	}

	for {
		// laço infinito
		fmt.Println("Loop infinito")
		time.Sleep(time.Second * 2)
	}
}
