package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func main() {
	// fmt.Println("Funções Panic e Recover")
	// funcao1()
	// funcao2()
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução")

}
