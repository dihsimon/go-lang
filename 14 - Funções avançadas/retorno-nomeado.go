package main

import (
	"fmt"
)

func calculoMatematico(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println("retorno nomeados")

	fmt.Println(calculoMatematico(10, 5))
}
