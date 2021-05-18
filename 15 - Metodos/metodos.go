package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar")
}

func main() {
	fmt.Println("Metodos")
	usuario1 := usuario{"Diego", 28}
	fmt.Println(usuario1)
	usuario1.salvar()
}
