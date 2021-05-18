package main

import "fmt"

func main() {
	canal := make(chan string, 200)

	canal <- "Olá mundo"
	canal <- "Programando em GO!"
	canal <- ""

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
