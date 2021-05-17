package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 int64 = 100000000000000
	fmt.Println(numero2)

	var real float32 = 1234.67
	fmt.Println(real)

	real2 := 234.56
	fmt.Println(real2)

	var str string = "texto"
	fmt.Println(str)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
