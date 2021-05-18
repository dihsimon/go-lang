package main

import (
	"fmt"
)

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funcções Variaticas")
	fmt.Println(soma(1, 4, 6, 8, 9, 4, 6, 7, 8))
	escrever("Olá Mundo", 1, 3, 4, 5, 6, 7, 8, 9)
}
