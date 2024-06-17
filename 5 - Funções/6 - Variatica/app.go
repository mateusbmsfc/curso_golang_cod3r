package main

func media(numeros  ...float64) float64 {
	soma := 0.0
    for _, numero := range numeros {
        soma += numero
    }
    return soma / float64(len(numeros))
}

func main() {
    media := media(1, 2, 3, 4, 5)
    println("Média: ", media)
}