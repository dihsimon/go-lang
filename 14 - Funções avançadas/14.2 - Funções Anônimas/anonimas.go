package main

import (
	"fmt"
)

func main() {
	fmt.Println("Funções anônimas")
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
