package main

import (
	"fmt"
)

func main() {
	var varialve1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(varialve1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel4"
		variavel4 string = "variavel5"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"

	fmt.Println(variavel5, variavel6)
}
