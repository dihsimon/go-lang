package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo de Structs")

	var u usuario
	u.nome = "Diego"
	u.idade = 28
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Luiz Scott", 10}

	usuario2 := usuario{"Diegão", 28, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Diegão2"}
	fmt.Println(usuario3)
}
