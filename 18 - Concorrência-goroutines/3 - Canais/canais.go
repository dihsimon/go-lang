package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escreverNaTela("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")

}
func escreverNaTela(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
