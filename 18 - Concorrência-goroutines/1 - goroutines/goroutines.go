package main

import (
	"fmt"
	"time"
)

func main() {
	go escreverNaTela("Olá Mundo")
	escreverNaTela("Programando em GO!")
}

func escreverNaTela(textto string) {
	for {
		fmt.Println(textto)
		time.Sleep(time.Second)
	}
}
