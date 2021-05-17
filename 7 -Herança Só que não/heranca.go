package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Diego", 28, 178, "simon"}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "Fatec Ourinhos"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.faculdade)
}
